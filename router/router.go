package router

import (
	"GoTwitter/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Mount() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	router.Route("/api", func(router chi.Router) {
		router.Get("/ping", handler.PingHandler)
	})

	return router
}
