package hlModel

type DeleteBannerModel struct {
	ID int
}

type DeleteBannerFeatureOrTagModel struct {
	FeatureID int `json:"feature_id"`
	TagID int `json:"tag_id"`
}
