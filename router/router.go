package router

import (
	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/handlers"
	"golang.org/x/net/websocket"
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

	// auth router
	profile := api.Group("/profile")
	profile.POST("/", handlers.Profile)
	profile.POST("/fullname", handlers.SetFullName)

	// auth router
	files := api.Group("/files")
	files.Any("/", handlers.FilesList)
	files.POST("/upload", handlers.UploadFile)
	files.POST("/newdir", handlers.NewDir)
	files.POST("/rename", handlers.RenameFile)
	files.POST("/copy", handlers.Copy)
	files.POST("/delete", handlers.DeleteFile)

	files.Any("/view", handlers.ViewFile)

	api.GET("/webterm", func(ctx *gin.Context) {
		hand := websocket.Handler(handlers.WebTerm)
		hand.ServeHTTP(ctx.Writer, ctx.Request)
	})

	api.GET("/proc", func(ctx *gin.Context) {
		hand := websocket.Handler(handlers.ProcMoni)
		hand.ServeHTTP(ctx.Writer, ctx.Request)
	})

	g.NoRoute(handlers.NotFound)
}
