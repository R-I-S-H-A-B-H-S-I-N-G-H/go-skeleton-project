package app

import (
	"hello/routes/hello"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	handleRoutes(router)
	return router
}

func handleRoutes(router chi.Router)  {
	router.Route("/hello", hello.HandleHelloRoutes)
}