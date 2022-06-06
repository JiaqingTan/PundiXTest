package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pundixtest/constant"
)

type Controller struct {}

func (controller *Controller) ExecuteCommand(ctx *gin.Context) {
	command1 := ctx.Param(constant.Command1)
	command2 := ctx.Param(constant.Command2)
	command3 := ctx.Param(constant.Command3)
	ctx.JSON(http.StatusOK, fmt.Sprintf("Command 1: %s, Command 2: %s, Command 3: %s", command1, command2, command3))
}