package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mostafa-DE/go-further-app/internal/data"
)

const version = "0.0.1"

type Config struct {
	port int
	env  string
}

type App struct {
	config Config
	logger *log.Logger
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.port, "port", 8000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "App envs (dev|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &App{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      // max time for connections using TCP Keep-Alive
		ReadTimeout:  10 * time.Second, // max time to read request from client
		WriteTimeout: 30 * time.Second, // max time to write response to client
	}

	logger.Printf("Starting server on port http://localhost:%d", cfg.port)
	err := server.ListenAndServe()

	if err != nil {
		logger.Fatalf("Error starting server: %v\n", err)
	}

}

func (app *App) healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	data := data.Envelope{
		"status": "OK",
		"system_info": data.SystemInfo{
			Version: version,
			Env:     app.config.env,
		},
	}

	err := app.writeJSON(res, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}
}
