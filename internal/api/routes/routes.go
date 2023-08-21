package routes

import (
	"github.com/franciscofferraz/go-struct/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/create", h.UserHandler.CreateUserHandler)
	})
}
