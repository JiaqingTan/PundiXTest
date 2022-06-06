package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {}

func (controller *Controller) HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}