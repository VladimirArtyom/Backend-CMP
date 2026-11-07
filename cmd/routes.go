package main

import (
	"compro-exh/internal"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"GET", "POST", "DELETE"},
			AllowedHeaders: []string{"Accept", "Content-Type"},
			ExposedHeaders: []string{"Link"},
			MaxAge:         300,
		}))

	mux.Post("/participant", internal.InsertParticipant)
	mux.Get("/participant", internal.GetParticipant)
	mux.Delete("/participant", internal.GetParticipant)

	return mux
}
