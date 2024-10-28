package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/health/", app.healthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id/", app.getMovie)
	router.HandlerFunc(http.MethodPost, "/v1/movies/", app.createNewMovie)

	return router
}
