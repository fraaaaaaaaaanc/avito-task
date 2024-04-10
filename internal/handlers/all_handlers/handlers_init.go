package allHandlers

import (
	mainStorage "avito-tech/internal/storage/main_storage"
	"avito-tech/internal/token"
	"github.com/go-playground/validator"
)

type Handlers struct {
	validator    *validator.Validate
	strg         *mainStorage.Storage
	tokenAccount *token.TokenAccount
}

func NewHandlers(storage *mainStorage.Storage, tokenAccount *token.TokenAccount) *Handlers {
	valid := validator.New()
	return &Handlers{
		strg:         storage,
		validator:    valid,
		tokenAccount: tokenAccount,
	}
}
