package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app:= &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	const PORT = "3000"
	server := &http.Server{
		Addr: ":"+PORT,
		Handler: a.router,
	}

	fmt.Println("Server starting at port :: ", PORT)

	err := server.ListenAndServe()

	if (err != nil){
		return fmt.Errorf("failed to start server :: %w", err);
	}
	return nil
}