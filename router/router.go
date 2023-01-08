package router

import (
	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/handlers"
)

func SetRoutes(g *gin.Engine) {

	// set middlewares
	g.Use(handlers.Cors)

	// home router
	home := g.Group("/")
	home.Any("/", handlers.Home)
	
	// api router
	api := g.Group("/api")
	api.Any("/", handlers.Api)

	// auth router
	auth := api.Group("/auth")
	auth.POST("/", handlers.Auth)

	g.NoRoute(handlers.NotFound)
}
