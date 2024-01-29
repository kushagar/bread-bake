package main

import (
	"bread-bake/pkg/config"
	"bread-bake/pkg/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(config *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
