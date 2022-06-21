package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheryarbutt/Learn-Go-Gin/entity"
	"github.com/sheryarbutt/Learn-Go-Gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {

	video := entity.Video{}
	ctx.BindJSON(&video)
	c.service.Save(&video)
	return video
}
