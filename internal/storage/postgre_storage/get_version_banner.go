package postgreStorage

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	"context"
)

func (ps *PostgreStorage) GetVersionBanner(ctx context.Context, bannerID int) (*[]hlModel.ResponseBannerModel, error) {
	return nil, nil
}
