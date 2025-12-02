package app

import (
	"GoTwitter/config/env"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func NewConfig() Config {
	addr := env.GetString("PORT", ":4001")

	return Config{
		Addr: addr,
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return server.ListenAndServe()
}
