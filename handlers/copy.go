package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type CopyRequest struct {
	Token   string `json:"token"`
	File    string `json:"file"`
	OldPath string `json:"old_path"`
	NewPath string `json:"new_path"`
	Cut     bool   `json:"cut"`
}

func Copy(c *gin.Context) {
	var req CopyRequest

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

	oldPath := user.HomeDir + req.OldPath + "/" + req.File
	newPath := user.HomeDir + req.NewPath + "/" + req.File
	fileLoca := newPath
	fileTmpNum := 1
	for {
		_, err = os.Stat(fileLoca)
		if err != nil {
			if os.IsNotExist(err) {
				// filenot exist
				break
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "internal server error",
				})
				return
			}
		} else {
			fileLoca = fmt.Sprintf("%s (%d)", newPath, fileTmpNum)
			fileTmpNum++
		}
	}

	if req.Cut {
		err = os.Rename(oldPath, fileLoca)
	} else {
		err = copyFile(oldPath, fileLoca)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "copy/cut folder successful",
	})
}

func copyFile(src, dest string) error {
	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		return errors.New("invalid")
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		return errors.New("invalid")
	}

	return nil
}
