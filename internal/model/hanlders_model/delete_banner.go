package hlModel

type DeleteBannerModel struct {
	Id    int `json:"id" validate:"required"`
	Token string
}
