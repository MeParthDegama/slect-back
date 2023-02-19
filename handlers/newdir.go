package handlers

import (
	"net/http"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type NewDirRequest struct {
	Token    string `json:"token"`
	DirName  string `json:"dir_name"`
	BasePath string `json:"base_path"`
}

func NewDir(c *gin.Context) {
	var req NewDirRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	username, err := utils.AuthToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "login error",
		})
		return
	}

	user, err := user.Lookup(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "login error",
		})
		return
	}

	err = os.Mkdir(user.HomeDir+req.BasePath+"/"+req.DirName, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "folder create successful",
	})
}
