package postgreStorage

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	"context"
)

func (ps *PostgreStorage) GetUserBanner(ctx context.Context, userBannerModel hlModel.GetUserBannerModel) (*hlModel.BannerContentModel, error) {
	return nil, nil
}
