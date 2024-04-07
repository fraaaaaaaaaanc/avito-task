package allHandlers

import (
	"avito-tech/internal/cookie"
	mainStorage "avito-tech/internal/storage/main_storage"
	"github.com/go-playground/validator"
)

type Handlers struct {
	validator *validator.Validate
	strg      *mainStorage.Storage
	cookie    *cookie.Cookie
}

func NewHandlers(storage *mainStorage.Storage, cookie *cookie.Cookie) *Handlers {
	valid := validator.New()
	return &Handlers{
		strg:      storage,
		validator: valid,
		cookie:    cookie,
	}
}
