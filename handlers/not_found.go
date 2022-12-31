package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {

	if strings.HasPrefix(c.Request.URL.Path, "/api/") {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "api not found",
		})

		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "path not found",
	})
}
