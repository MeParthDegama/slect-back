package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/config"
	"github.com/parthkax70/slect/utils"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(c *gin.Context) {

	var info LoginInfo

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "only accept json body",
		})
		return
	}

	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	if info.Password == "" || info.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid request field",
		})
		return
	}

	loginSucc := utils.CheckSystemPassword(info.Username, info.Password)

	var res map[string]any

	if loginSucc {
		
		token := utils.GenToken()
		
		config.AddToken(token, info.Username)

		res = gin.H{
			"status":  true,
			"message": "login successful",
			"token": token,
		}
	} else {
		res = gin.H{
			"status":  false,
			"message": "Invalid username or password",
		}
	}

	c.JSON(http.StatusOK, res)
}
