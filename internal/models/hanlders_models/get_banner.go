package hlModel

type GetBannerModel struct {
	TagId     int `schema:"tag_id,omitempty"`
	FeatureId int `schema:"feature_id,omitempty"`
	Limit     int `schema:"limit,omitempty"`
	Offset    int `schema:"offset,omitempty"`
}

type ResponseBannerModel struct {
	BannerID  int                `json:"banner_id"`
	TagIDs    []int              `json:"tag_ids"`
	FeatureID int                `json:"feature_id"`
	IsActive  bool               `json:"is_active"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Content   BannerContentModel `json:"content"`
}
