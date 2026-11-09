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

	mux.Get("/group", internal.GetGroup)
	mux.Post("/group", internal.InsertGroup)
	mux.Delete("/group", internal.DeleteGroup)
	mux.Get("/groups", internal.GetGroups)

	mux.Post("/vote", internal.InsertVote)
	mux.Delete("/vote", internal.DeleteVote)
	mux.Get("/vote", internal.GetVote)
	mux.Get("/votes", internal.GetVotes)

	return mux
}
