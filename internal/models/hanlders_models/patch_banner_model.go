package hlModel

type PatchBannerModel struct {
	TagIds    []int              `json:"tag_ids,required"`
	FeatureId int                `json:"feature_id,required"`
	Content   BannerContentModel `json:"content,required"`
	IsActive  bool               `json:"is_active,required"`
	Id        uint32
}
