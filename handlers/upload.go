package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

func UploadFile(c *gin.Context) {

	token := c.Request.FormValue("token")
	basePath := c.Request.FormValue("base_path")
	if token == "" || basePath == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	username, err := utils.AuthToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "login error",
		})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
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

	fileSplitStr := strings.Split(user.HomeDir+basePath+"/"+fileHeader.Filename, ".")
	fileNameWithOutExt := strings.Join(fileSplitStr[:len(fileSplitStr)-1], ".")
	fileExt := fileSplitStr[len(fileSplitStr)-1]
	fileLoca := fmt.Sprintf("%s.%s", fileNameWithOutExt, fileExt)
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
			fileLoca = fmt.Sprintf("%s (%d).%s", fileNameWithOutExt, fileTmpNum, fileExt)
			fileTmpNum++
		}
	}

	err = os.WriteFile(fileLoca, fileBytes, 0644)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Invalid request",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "File upload successfull",
	})
}
