package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
)

var Cors = cors.New(cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "PATCH"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return true
	},
	MaxAge: 74 * time.Hour,
})
