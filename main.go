package main

import (
	"GoTwitter/app"
	"GoTwitter/config/env"
)

func main() {
	env.Load()

	cfg := app.NewConfig()
	aap := app.NewApplication(cfg)

	aap.Run()

}
