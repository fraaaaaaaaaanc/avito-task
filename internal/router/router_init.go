package router

import (
	"avito-tech/internal/cookie"
	allHandlers "avito-tech/internal/handlers/all_handlers"
	"github.com/go-chi/chi"
)

func NewRouter(hndlr *allHandlers.Handlers, cookies *cookie.Cookie) (chi.Router, error) {
	r := chi.NewRouter()

	//r.Use(compress.MiddlewareCompress(),
	//	logger.MiddlewareHandlerLog())

	//r.Route("/api/user/", func(r chi.Router) {
	//	r.Route("/orders", func(r chi.Router) {
	//		r.With(cookies.MiddlewareCheckCookie()).Get("/", hndlr.GetOrders)
	//		r.With(cookies.MiddlewareCheckCookie()).Post("/", hndlr.PostOrders)
	//	})
	//
	//	r.With(cookies.MiddlewareCheckCookie()).Get("/balance", hndlr.Balance)
	//	r.With(cookies.MiddlewareCheckCookie()).Get("/withdrawals", hndlr.Withdrawals)
	//
	//	r.Post("/register", hndlr.Register)
	//	r.Post("/login", hndlr.Login)
	//	r.With(cookies.MiddlewareCheckCookie()).Post("/balance/withdraw", hndlr.WithDraw)
	//})

	return r, nil
}
