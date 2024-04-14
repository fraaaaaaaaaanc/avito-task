package hlModel

type GetUserBannerModel struct {
	TagId           int  `schema:"tag_id,required"`
	FeatureId       int  `schema:"feature_id,required"`
	UseLastRevision bool `schema:"use_last_revision,omitempty"`
}
