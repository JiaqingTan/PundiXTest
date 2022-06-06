package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pundixtest/constant"
	"pundixtest/controller"
)

type Router struct {
	engine *gin.Engine
}

func (router *Router) Init() {
	router.engine = new(gin.Engine)
	router.engine = gin.Default()
}

func (router *Router) ServeRoutes(controller *controller.Controller) {
	router.engine.GET(fmt.Sprintf("/:%s/:%s/:%s", constant.Command1, constant.Command2, constant.Command3), controller.ExecuteCommand)
}

// Run on localhost by default
func (router *Router) Run(port int) {
	router.engine.Run(fmt.Sprintf(":%d", port))
}