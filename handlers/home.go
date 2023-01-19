package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/config"
)

func Home(c *gin.Context) {

	fmt.Println(config.AppConfig) // tmp

	c.JSON(http.StatusOK, gin.H{
		"message": "slect web media",
	})
}
