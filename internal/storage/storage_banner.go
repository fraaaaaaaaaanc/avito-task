package storage

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	storageModels "avito-tech/internal/models/storage_model"
	"context"
)

type StorageBanner interface {
	GetUserBanner(context.Context, hlModel.GetUserBannerModel) (*hlModel.BannerContentModel, error)
	GetBanner(context.Context, hlModel.GetBannerModel) (string, error)
	GetVersionBanner(context.Context, int) (*[]hlModel.ResponseBannerModel, error)
	SetBanner(context.Context, *hlModel.PostBannerModel) (uint32, error)
	PatchBanner(context.Context, *hlModel.PatchBannerModel) error
	DelBanner(context.Context, int) error
	DelBannerFeatureOrTag(context.Context, []storageModels.DelFeatureOrTagChan) error
	CloseDB()
}
