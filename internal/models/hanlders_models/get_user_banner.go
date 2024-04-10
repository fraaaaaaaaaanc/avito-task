package hlModel

type GetUserBannerModel struct {
	TagId           int  `schema:"tag_id" validate:"required"`
	FeatureId       int  `schema:"feature_id" validate:"required"`
	UseLastRevision bool `schema:"use_last_revision,omitempty"`
}
