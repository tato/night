package main

import (
	"github.com/gin-gonic/gin"

	"supok.es/night/basepath"
	"supok.es/night/files"
)

func main() {
	router := gin.Default()

	router.GET("/basepath", basepath.GetBasePath)
	router.PUT("/basepath", basepath.PutBasePath)

	router.GET("/files", files.GetFileTree)
	router.GET("/files/*path", files.GetFile)

	router.Run("localhost:9001")
}
