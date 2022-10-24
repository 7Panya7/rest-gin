package video

import "github.com/gin-gonic/gin"

type VideoHandler interface {
	FindAll() []Video
	Save(ctx *gin.Context) Video
}
