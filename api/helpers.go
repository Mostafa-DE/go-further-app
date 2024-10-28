package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *App) parseParam(r *http.Request, key string) string {
	params := httprouter.ParamsFromContext(r.Context())

	return params.ByName(key)
}

func (app *App) strToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}

	return int(i)
}

func (app *App) writeJSON(res http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	_res, err := json.Marshal(data)

	if err != nil {
		app.logger.Println(err)

		return err
	}

	for key, val := range headers {
		res.Header()[key] = val
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(_res)

	return nil
}
