package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type application struct {
	config config
}

type config struct {
	addr   string
	env    string
	apiURL string
}

func (app *application) serve() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", app.healthCheckHandler)
	r.Post("/create", app.createShortURL)

	return r

}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:              app.config.addr,
		Handler:           mux,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}
