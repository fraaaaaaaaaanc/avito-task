package hlModel

type PostBannerModel struct {
	TagIds    []int              `json:"tag_ids,omitempty"`
	FeatureId int                `json:"feature_id,omitempty"`
	IsActive  bool               `json:"is_active,omitempty"`
	Content   BannerContentModel `json:"content,omitempty"`
}
