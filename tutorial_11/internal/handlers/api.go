package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	// our custom packages
	"tutorial_11/internal/middleware"
)

// functions that start with upper case letters are public function
// functions that start with lower case letters are private function
func Handler(r *chi.Mux) {
	// global middleware to handle stripping slashes at the end of endpoint
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization) // our custom middleware

		router.Get("/coins", GetCoinBalance) // HTTP GET endpoint
	})
}
