package router

import (
	"avito-tech/internal/caches"
	allHandlers "avito-tech/internal/handlers/all_handlers"
	"avito-tech/internal/token"

	"github.com/go-chi/chi"
)

func NewRouter(hndlr *allHandlers.Handlers, token *token.TokenAccount, redis *caches.Cache) (chi.Router, error) {
	r := chi.NewRouter()

	r.Route("/user_banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Get("/", hndlr.GetUserBannerHandler)
	})

	r.Route("/banner", func(r chi.Router) {
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Get("/", hndlr.GetBannerHandler)
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Post("/", hndlr.PostBannerHandler)
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Patch("/{id:[0-9]+}", hndlr.PatchBannerIdHandler)
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Delete("/{id:[0-9]+}", hndlr.DeleteBannerIdHandler)
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Get("/versions/{id:[0-9]+}", hndlr.GetVersionBannerHandler)
		r.With(token.MiddlewareCheckToken(), redis.CacheMiddleware).Delete("/featueortag/{id:[0-9]+}", hndlr.GetVersionBannerHandler)
	})

	r.Route("/login", func(r chi.Router) {
		r.Get("/", hndlr.LoginHandler)
	})

	return r, nil
}
