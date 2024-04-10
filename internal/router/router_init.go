package router

import (
	allHandlers "avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/token"
	"github.com/go-chi/chi"
)

func NewRouter(hndlr *allHandlers.Handlers, token *token.TokenAccount) (chi.Router, error) {
	r := chi.NewRouter()

	r.Route("/user_banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken()).Get("/", hndlr.GetUserBanner)
	})

	r.Route("/banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken()).Get("/", hndlr.GetBanner)
		r.With(token.MiddlewareCheckToken()).Post("/", hndlr.PostBanner)
		r.With(token.MiddlewareCheckToken()).Patch("/{id:[0-9]+}", hndlr.PatchBannerId)
		r.With(token.MiddlewareCheckToken()).Delete("/{id:[0-9]+}", hndlr.DeleteBannerId)
	})

	r.Route("/login", func(r chi.Router) {
		r.Get("/", hndlr.Login)
	})

	return r, nil
}
