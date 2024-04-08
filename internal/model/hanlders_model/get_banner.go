package hlModel

type GetBannerModel struct {
	TagId     int `schema:"tag_id,omitempty"`
	FeatureId int `schema:"feature_id,omitempty"`
	Limit     int `schema:"limit,omitempty"`
	Offset    int `schema:"offset,omitempty"`
	Token     string
}
