package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"supok.es/night/basepath"
	"supok.es/night/files"
)

func main() {

	opt := parseArguments()

	if !opt.debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	if opt.debugMode {
		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		}))
	}

	router.GET("/basepath", basepath.GetBasePath)
	router.PUT("/basepath", basepath.PutBasePath)

	router.GET("/files", files.GetFileTree)
	router.GET("/files/*path", files.GetFile)

	router.Run("localhost:9001")
}

type options struct {
	debugMode bool
}

func parseArguments() options {
	var opt options

	for _, arg := range os.Args[1:] {
		if arg == "--debug" {
			opt.debugMode = true
		}
	}

	return opt
}
