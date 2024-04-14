package hlModel

type PostBannerModel struct {
	TagIds    []uint32              `json:"tag_ids,required"`
	FeatureId int                `json:"feature_id,required"`
	IsActive  bool               `json:"is_active,required"`
	Content   BannerContentModel `json:"content,required"`
}

type ResponsePostBannerModel struct {
	BannerID int `json:"banner_id"`
}
