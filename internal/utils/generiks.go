package utils

import hlModel "avito-tech/internal/models/hanlders_models"

type Responses interface {
	*hlModel.BannerContentModel | hlModel.ResponseBannerModel | hlModel.ResponsePostBannerModel
}
