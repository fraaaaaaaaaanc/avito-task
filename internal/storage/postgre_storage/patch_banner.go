package postgreStorage

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"context"
	"encoding/json"
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func updateBannerInfo(tx *sql.Tx, id uint32, bannerModel *hlModel.PatchBannerModel) error {
	switch {
	// Если у нас изменился только айди фичи, её можно обновить по id баннера
	case bannerModel.TagIDs != nil && bannerModel.FeatureID != 0:
		if _, err := tx.Exec(updateFeaturesQuery, id, bannerModel.FeatureID); err != nil {
			return errors.Wrapf(checkPgConflictError(err),
				"can't update feature id %d to banner", bannerModel.FeatureID)
		}
	// Если у нас изменился список тэгов, то нужно сначала удалить все записи с тэгами, а потом их снова создать
	case bannerModel.TagIDs != nil:
		var featureID uint32
		if err := tx.QueryRow(deleteFeaturesTagsQuery, id).Scan(&featureID); err != nil {
			return errors.Wrap(err, "can't delete feature id and tag ids of banner")
		}

		if bannerModel.FeatureID != 0 {
			featureID = bannerModel.FeatureID
		}

		if _, err := tx.Exec(addFeaturesAndTagsQuery, id, featureID,
			pgtype.FlatArray[uint32](bannerModel.TagIDs)); err != nil {
			return errors.Wrapf(checkPgConflictError(err),
				"can't add feature id %d and tag ids %v to banner", featureID, bannerModel.TagIDs)
		}
	}

	return nil
}


func (ps *PostgreStorage) PatchBanner(ctx context.Context, bannerModel *hlModel.PatchBannerModel) error {
	var updatedID uint32

	if err := WithTransaction(ps.db,
		func(tx *sql.Tx) error {
			if err := tx.QueryRow(checkDeleted, bannerModel.Id).Scan(&updatedID); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return ErrorBannerNotFound
				}

				return errors.Wrapf(err, "can't check banner on deleted")
			}

			if bannerModel.IsActive {
				if err := tx.QueryRow(updateActiveQuery,
					bannerModel.Id, bannerModel.IsActive).
					Scan(&updatedID); err != nil {
					if errors.Is(err, pgx.ErrNoRows) {
						return ErrorBannerNotFound
					}

					return errors.Wrapf(err, "can't update banner")
				}
			}

			content, err := json.Marshal(bannerModel.Content)
			if err != nil {
				logger.Error("Create content", zap.Error(err))
				return err
			}

			if content != nil {
				if err := addContent(tx, bannerModel.Id, string(content)); err != nil {
					return err
				}
			}

			return updateBannerInfo(tx, bannerModel.Id, bannerModel)
		},
	); err != nil {
		return errors.Wrapf(err, "when updating banner with id %d", bannerModel.Id)
	}

	return nil
}
