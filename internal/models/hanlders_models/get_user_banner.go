package hlModel

type GetUserBannerModel struct {
	TagId           int  `schema:"tag_id"`
	FeatureId       int  `schema:"feature_id"`
	UseLastRevision bool `schema:"use_last_revision,omitempty"`
}
