package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sheryarbutt/Learn-Go-Gin/controller"
	"github.com/sheryarbutt/Learn-Go-Gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {

	server := gin.Default()
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
