package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mostafa-DE/go-further-app/internal/data"
)

func (app *App) createNewMovie(res http.ResponseWriter, req *http.Request) {
	var Body struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(res, req, &Body)

	if err != nil {
		app.badRequest(res, req, err)
		return
	}

	fmt.Fprintf(res, "%+v\n", Body)
}

func (app *App) getMovie(res http.ResponseWriter, req *http.Request) {
	id := app.strToInt(app.parseParam(req, "id"))

	if id <= 0 {
		app.notFound(res, req)
		return
	}

	movie := data.Movie{
		ID:         int64(id),
		Created_at: time.Now(),
		Title:      "Awesome movie",
		Year:       int32(time.Now().Year()),
		Runtime:    120,
		Genres:     []string{"Action", "Adventure"},
		Version:    1,
	}

	err := app.writeJSON(res, http.StatusOK, data.Envelope{"movie": movie}, nil)

	if err != nil {
		app.internalServerError(res, req, err)
	}
}

func (app *App) getAllMovies(res http.ResponseWriter, req *http.Request) {
	_res := map[string]string{
		"message": "All movies",
	}

	err := app.writeJSON(res, http.StatusOK, data.Envelope{"movies": _res}, nil)

	if err != nil {
		app.internalServerError(res, req, err)
	}
}
