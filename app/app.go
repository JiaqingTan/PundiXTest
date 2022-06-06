package app

import (
	"pundixtest/config"
	"pundixtest/controller"
	"pundixtest/router"
)

type App struct {
	config *config.AppConfig
	router router.Router
}

func New() *App {
	return new(App)
}

func (app *App) Init(appConfig *config.AppConfig, controller *controller.Controller) {
	app.config = appConfig
	app.router.Init()
	app.router.ServeRoutes(appConfig, controller)
}

func (app *App) Start() {
	app.router.Run(app.config.Port)
}