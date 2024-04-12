package router

import (
	allHandlers "avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/token"
	"github.com/go-chi/chi"
)

func NewRouter(hndlr *allHandlers.Handlers, token *token.TokenAccount) (chi.Router, error) {
	r := chi.NewRouter()

	r.Route("/user_banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken()).Get("/", hndlr.GetUserBannerHandler)
	})

	r.Route("/banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken()).Get("/", hndlr.GetBannerHandler)
		r.With(token.MiddlewareCheckToken()).Post("/", hndlr.PostBannerHandler)
		r.With(token.MiddlewareCheckToken()).Patch("/{id:[0-9]+}", hndlr.PatchBannerIdHandler)
		r.With(token.MiddlewareCheckToken()).Delete("/{id:[0-9]+}", hndlr.DeleteBannerIdHandler)
		r.With(token.MiddlewareCheckToken()).Get("/versions/{id:[0-9]+}", hndlr.GetVersionBannerHandler)
	})

	r.Route("/login", func(r chi.Router) {
		r.Get("/", hndlr.LoginHandler)
	})

	return r, nil
}
