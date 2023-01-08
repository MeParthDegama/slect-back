package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(c *gin.Context) {

	var info LoginInfo

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "only accept json body",
		})
		return
	}

	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request",
		})
		return
	}

	if info.Password == "" || info.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request field",
		})
		return
	}

	loginSucc := utils.CheckSystemPassword(info.Username, info.Password)

	var res map[string]any

	if loginSucc {
		res = gin.H{
			"status":  true,
			"message": "login successful",
		}
	} else {
		res = gin.H{
			"status":  false,
			"message": "login error",
		}
	}

	c.JSON(http.StatusOK, res)
}
