package hello

import (
	"hello/controllers/hello"

	"github.com/go-chi/chi/v5"
)

func HandleHelloRoutes(router chi.Router)  {
	router.Get("/{times}", hello.Get)
}
