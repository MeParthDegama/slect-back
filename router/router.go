package router

import (
	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/handlers"
)

func SetRoutes(g *gin.Engine) {

	// home router
	home := g.Group("/")
	home.Any("/", handlers.Home)
	
	// api router
	api := g.Group("/api")
	api.Any("/", handlers.Auth)

	g.NoRoute(handlers.NotFound)
}
