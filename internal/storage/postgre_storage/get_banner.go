package postgreStorage

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	"github.com/jackc/pgx/v5"
	"context"
	"github.com/pkg/errors"
)

func (ps *PostgreStorage) GetBanner(ctx context.Context, bannerModel hlModel.GetBannerModel) (string, error) {
	var content string
	if err := ps.db.QueryRow(getQuery, bannerModel.FeatureId, bannerModel.TagId, 1).
		Scan(
			&content,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return content, errors.Wrapf(ErrorBannerNotFound,
				"with feature id %d and tag id %d and version %v", bannerModel.FeatureId, bannerModel.TagId, 1)
		}

		return content, errors.Wrapf(err,
			"can't get banner with feature id %d and tag id %d and version %v", bannerModel.FeatureId, bannerModel.TagId, 1)
	}

	return content, nil
}
