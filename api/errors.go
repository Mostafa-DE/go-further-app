package main

import (
	"fmt"
	"net/http"

	"github.com/Mostafa-DE/go-further-app/internal/data"
)

func (app *App) logError(req *http.Request, err error) {
	app.logger.Println(err)
}

func (app *App) errorResponse(res http.ResponseWriter, req *http.Request, status int, msg string) {
	env := data.Envelope{"error": msg}
	err := app.writeJSON(res, status, env, nil)

	if err != nil {
		app.logError(req, err)
		http.Error(res, "Internal Server Error", 500)
	}
}

func (app *App) internalServerError(res http.ResponseWriter, req *http.Request, err error) {
	app.logError(req, err)

	msg := "Something went wrong, please contact the support team for help."
	app.errorResponse(res, req, http.StatusInternalServerError, msg)
}

func (app *App) notFound(res http.ResponseWriter, req *http.Request) {
	app.logError(req, nil)

	msg := "The requested resource could not be found."
	app.errorResponse(res, req, http.StatusNotFound, msg)
}

func (app *App) methodNotAllowed(res http.ResponseWriter, req *http.Request) {
	app.logError(req, nil)

	msg := fmt.Sprintf("The %s method is not supported for this resource.", req.Method)
	app.errorResponse(res, req, http.StatusMethodNotAllowed, msg)
}

func (app *App) badRequest(res http.ResponseWriter, req *http.Request, err error) {
	app.logError(req, err)

	app.errorResponse(res, req, http.StatusBadRequest, err.Error())
}
