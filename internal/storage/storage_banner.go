package storage

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	storageModels "avito-tech/internal/models/storage_model"
	"context"
)

type StorageBanner interface {
	GetUserBanner(context.Context, hlModel.GetUserBannerModel) (*hlModel.BannerContentModel, error)
	GetBanner(context.Context, hlModel.GetBannerModel) (*hlModel.ResponseBannerModel, error)
	GetVersionBanner(context.Context, int) (*[]hlModel.ResponseBannerModel, error)
	SetBanner(context.Context, hlModel.PostBannerModel) (*hlModel.ResponsePostBannerModel, error)
	PatchBanner(context.Context, *hlModel.PatchBannerModel) error
	DelBanner(context.Context, int) error
	DelBannerFeatureOrTag(ctx context.Context, tagChan []storageModels.DelFeatureOrTagChan) error
}
