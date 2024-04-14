package postgreStorage

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"github.com/jackc/pgx/v5/pgtype"
)

func addContent(tx *sql.Tx, id uint32, content string) error {
	if _, err := tx.Exec(addContentQuery, id, content); err != nil {
		return errors.Wrap(err, "can't add content to banner")
	}

	return nil
}

func (ps *PostgreStorage) SetBanner(ctx context.Context, bannerModel *hlModel.PostBannerModel) (uint32, error) {
	var createdID uint32

	if err := WithTransaction(ps.db,
		func(tx *sql.Tx) error {
			if err := tx.QueryRow(createQuery, bannerModel.IsActive).
				Scan(
					&createdID,
				); err != nil {
				return errors.Wrap(err, "can't create banner")
			}

			content, err := json.Marshal(bannerModel.Content)
			if err != nil {
				logger.Error("error create conten", zap.Error(err))
				return err
			}

			if err := addContent(tx, createdID, string(content)); err != nil {
				return err
			}

			if _, err := tx.Exec(addFeaturesAndTagsQuery, createdID, bannerModel.FeatureId,
				pgtype.FlatArray[uint32](bannerModel.TagIds)); err != nil {
				return errors.Wrapf(checkPgConflictError(err),
					"can't add feature id %d and tag ids %v to banner", bannerModel.FeatureId, bannerModel.TagIds)
			}

			return nil
		},
	); err != nil {
		return 0, errors.Wrap(err, "when creating banner")
	}

	return createdID, nil
}
