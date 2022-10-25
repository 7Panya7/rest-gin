package main

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
	"rest-gin/internal/auth"
	"rest-gin/internal/video"
	"rest-gin/pkg/db/postgres"
	"rest-gin/pkg/logger"
)
import _ "github.com/lib/pq"

var (
	vServ = video.NewService()
	vHand = video.NewHandler(vServ)
)

func setupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {

	setupLogOutput()

	postgres.GetDB()

	server := gin.New()

	server.Use(gin.Recovery(), logger.Logger(), auth.Auth(), gindump.Dump())

	server.GET("/video", func(ctx *gin.Context) {
		ctx.JSON(200, vHand.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		err := vHand.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Video Input is valid!"})
		}
	})

	server.Run(":8089")
}
