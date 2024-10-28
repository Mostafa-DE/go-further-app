package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mostafa-DE/go-further-app/internal/data"
)

func (app *App) createNewMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create new movie")
}

func (app *App) getMovie(res http.ResponseWriter, req *http.Request) {
	id := app.strToInt(app.parseParam(req, "id"))

	if id <= 0 {
		app.logger.Printf("Invalid movie id: %d", id)
		http.Error(res, "Invalid movie id", http.StatusBadRequest)

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
		app.logger.Println(err)
		http.Error(res, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *App) getAllMovies(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all movies")

	res.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"message": "All movies",
	}

	err := json.NewEncoder(res).Encode(data)

	if err != nil {
		app.logger.Println(err)
	}
}
