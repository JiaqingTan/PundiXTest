package main

import (
	"log"
	"pundixtest/app"
	"pundixtest/config"
	"pundixtest/controller"
)

func main() {
	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Println("App cannot be initialized")
		return
	}

	controller := new(controller.Controller)
	controller.Init(appConfig)
	if err != nil {
		log.Println("App cannot be initialized")
		return
	}

	app := app.New()
	app.Init(appConfig, controller)
	app.Start()
}

