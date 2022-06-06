package main

import (
	"log"
	app "pundixtest/app"
	"pundixtest/config"
	"pundixtest/controller"
)

func main() {

	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Println("App cannot be initialized")
	}

	controller := new(controller.Controller)

	app := app.New()
	app.Init(appConfig, controller)
	app.Start()
}

