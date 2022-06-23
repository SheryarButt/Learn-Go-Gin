package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sheryarbutt/Learn-Go-Gin/controller"
	"github.com/sheryarbutt/Learn-Go-Gin/middlewares"
	"github.com/sheryarbutt/Learn-Go-Gin/repository"
	"github.com/sheryarbutt/Learn-Go-Gin/service"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOuput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOuput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth()) //, gindump.Dump())

	apiRoputes := server.Group("/api")
	{
		apiRoputes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})

		apiRoputes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video saved successfully",
				})
			}
		})

		apiRoputes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video updated successfully",
				})
			}
		})

		apiRoputes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video deleted successfully",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := server.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
