package handlers

import (
	"github.com/Obixander/GoApi/internal/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Autherization)

		router.Get("/coins", GetCoinBalance)
	})
}
