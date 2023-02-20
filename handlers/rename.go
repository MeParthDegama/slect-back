package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type RenameFileRequest struct {
	Token       string `json:"token"`
	OldFileName string `json:"old_file_name"`
	NewFileName string `json:"new_file_name"`
	BasePath    string `json:"base_path"`
}

func RenameFile(c *gin.Context) {
	var req RenameFileRequest

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

	fmt.Println(user.HomeDir, req.BasePath, req.NewFileName, req.OldFileName)

	basePath := user.HomeDir + req.BasePath

	err = os.Rename(basePath+"/"+req.OldFileName, basePath+"/"+req.NewFileName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "rename folder successful",
	})
}
