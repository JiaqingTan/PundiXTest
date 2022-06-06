package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pundixtest/config"
	"pundixtest/constant"
	"pundixtest/util"
)

type Controller struct {}

func (controller *Controller) ExecuteCommand(ctx *gin.Context) {
	command1 := ctx.Param(constant.Command1)
	command2 := ctx.Param(constant.Command2)
	command3 := ctx.Param(constant.Command3)
	ctx.JSON(http.StatusOK, fmt.Sprintf("Command 1: %s, Command 2: %s, Command 3: %s", command1, command2, command3))
}

func (controller *Controller) ValidateExecuteCommandParams(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		command1Param := ctx.Param(constant.Command1)
		command2Param := ctx.Param(constant.Command2)

		if !util.SliceContains(appConfig.AllowedCommands[constant.Command1], command1Param) ||
			!util.SliceContains(appConfig.AllowedCommands[constant.Command2], command2Param) {
			controller.NotFound(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (controller *Controller) NotFound(ctx *gin.Context) {
	ctx.JSON(404, "404 Not Found")
}