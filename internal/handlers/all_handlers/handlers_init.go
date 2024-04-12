package allHandlers

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	"avito-tech/internal/storage"
	"avito-tech/internal/token"
	"encoding/json"
)

type Handlers struct {
	strg                storage.StorageBanner
	tokenAccount        *token.TokenAccount
	errorIncorrectData  []byte
	errorInternalServer []byte
}

func NewHandlers(storage storage.StorageBanner, tokenAccount *token.TokenAccount) (*Handlers, error) {
	errorIncorrectData, err := json.Marshal(hlModel.ErrorBannerModel{Error: "incorrect data"})
	if err != nil {
		return nil, err
	}
	errorInternalServer, err := json.Marshal(hlModel.ErrorBannerModel{Error: "internal server error"})
	if err != nil {
		return nil, err
	}
	return &Handlers{
		strg:                storage,
		tokenAccount:        tokenAccount,
		errorIncorrectData:  errorIncorrectData,
		errorInternalServer: errorInternalServer,
	}, nil
}
