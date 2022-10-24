package video

import "github.com/gin-gonic/gin"

type handler struct {
	service VideoService
}

func NewHandler(service VideoService) VideoHandler {
	return &handler{
		service: service,
	}
}

func (h *handler) FindAll() []Video {
	return h.service.FindAll()
}
func (h *handler) Save(ctx *gin.Context) Video {
	var video Video
	ctx.BindJSON(&video)
	h.service.Save(video)
	return video
}
