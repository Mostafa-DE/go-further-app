package data

import "time"

type Envelope map[string]interface{}

type SystemInfo struct {
	Version string `json:"version"`
	Env     string `json:"env"`
}

type Movie struct {
	ID         int64     `json:"id"`
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Year       int32     `json:"year,omitempty"`
	Runtime    int32     `json:"runtime,omitempty,string"`
	Genres     []string  `json:"genres,omitempty"`
	Version    int32     `json:"version"`
}
