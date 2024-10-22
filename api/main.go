package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := App{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/health/", app.healthCheckHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,      // max time for connections using TCP Keep-Alive
		ReadTimeout:  10 * time.Second, // max time to read request from client
		WriteTimeout: 30 * time.Second, // max time to write response to client
	}

	logger.Printf("Starting server on port http://localhost:%d/v1/health/", cfg.port)

	err := server.ListenAndServe()

	if err != nil {
		logger.Fatalf("Error starting server: %v\n", err)
	}

}

func (app *App) healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"version": version,
		"env":     app.config.env,
		"status":  "OK",
	}

	err := json.NewEncoder(res).Encode(data)

	if err != nil {
		app.logger.Printf("Error encoding JSON: %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

}
