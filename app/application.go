package app

import (
	"GoTwitter/config/db"
	"GoTwitter/config/env"
	"GoTwitter/router"
	"fmt"
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

	_, dbErr := db.SetupDB()
	if dbErr != nil {
		fmt.Println("DB connection error", dbErr)
		return dbErr
	}

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.Mount(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return server.ListenAndServe()
}
