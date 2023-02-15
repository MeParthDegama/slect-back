package handlers

import (
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type FilesListRequest struct {
	Token string `json:"token"`
	Path  string `json:"path"`
}

type FilesInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isdir"`
	Size  int64  `json:"size"`
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

	reqDir := user.HomeDir + req.Path

	files, err := ioutil.ReadDir(reqDir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "dir not found",
		})
		return
	}

	filesListInfo := []FilesInfo{}

	for _, file := range files {
		filesListInfo = append(filesListInfo, FilesInfo{Name: file.Name(), IsDir: file.IsDir(), Size: file.Size()})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"message":  "success",
		"filelist": filesListInfo,
	})

}
