package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homepage(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func main() {
	router := gin.Default()
	router.GET("/", homepage)
	router.Run("localhost:3000")
}

