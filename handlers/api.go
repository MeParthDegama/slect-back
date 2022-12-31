package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Api(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "slect web media interface api",
	})
}
