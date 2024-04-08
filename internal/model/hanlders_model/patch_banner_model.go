package hlModel

type PatchBannerModel struct {
	TagIds    []int              `json:"tag_ids,omitempty"`
	FeatureId int                `json:"feature_id,omitempty"`
	Content   BannerContentModel `json:"content,omitempty"`
	IsActive  bool               `json:"is_active,omitempty"`
	Token     string
	Id        int
}
