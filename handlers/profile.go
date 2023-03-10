package handlers

import (
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type ProfileRequest struct {
	Token string `json:"token"`
}

func Profile(c *gin.Context) {

	var req ProfileRequest

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

	me, err := user.Lookup(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "login error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"username": username,
		"fullname": me.Name,
		"message":  "slect web media",
	})
}

type SetFullNameRequest struct {
	Token    string `json:"token"`
	FullName string `json:"fullname"` // new fullname
}

func SetFullName(c *gin.Context) {

	var req SetFullNameRequest

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

	if len(req.FullName) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	utils.SetFullName(username, req.FullName)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "fullname set successful",
	})
}
