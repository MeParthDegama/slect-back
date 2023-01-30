package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type FilesListRequest struct {
	Token string `json:"token"`
}

func FilesList(c *gin.Context) {

	var req FilesListRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	_, err = utils.AuthToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "login error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "success",
	})

}