package main

import (
	"github.com/gin-gonic/gin"
	"rest-gin/internal/video"
)

var (
	vServ video.VideoService = video.NewService()
	vHand video.VideoHandler = video.NewHandler(vServ)
)

func main() {
	server := gin.Default()

	server.GET("/video", func(ctx *gin.Context) {
		ctx.JSON(200, vHand.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(201, vHand.Save(ctx))
	})

	server.Run(":8081")
}
