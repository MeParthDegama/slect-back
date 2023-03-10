package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/config"
	"github.com/parthkax70/slect/router"
)

func main() {

	config.MakeConfigFile()
	config.LoadConfig()

	r := gin.New()

	router.SetRoutes(r)

	r.Run()

}
