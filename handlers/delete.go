package handlers

import (
	"net/http"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type DeleteFileRequest struct {
	Token     string `json:"token"`
	FileName  string `json:"file_name"`
	BasePath  string `json:"base_path"`
	Permanent bool   `json:"permanent"`
}

func DeleteFile(c *gin.Context) {
	var req DeleteFileRequest

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

	_, err = os.Stat(user.HomeDir + "/.delete/")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(user.HomeDir+"/.delete/", os.ModePerm)
		}
	}

	if req.Permanent {
		err = os.RemoveAll(user.HomeDir + req.BasePath + "/" + req.FileName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "internal server error",
			})
			return
		}
	} else {
		err = os.Rename(user.HomeDir+req.BasePath+"/"+req.FileName, user.HomeDir+"/.delete/"+req.FileName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "internal server error",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "rename folder successful",
	})
}
