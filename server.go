package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sheryarbutt/Learn-Go-Gin/controller"
	"github.com/sheryarbutt/Learn-Go-Gin/middlewares"
	"github.com/sheryarbutt/Learn-Go-Gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOuput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOuput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
