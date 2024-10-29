package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *App) readJSON(res http.ResponseWriter, req *http.Request, dist interface{}) error {
	err := json.NewDecoder(req.Body).Decode(dist)

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError

	if err != nil {
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}

			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return fmt.Errorf("Request body must not be empty")

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	return nil
}
