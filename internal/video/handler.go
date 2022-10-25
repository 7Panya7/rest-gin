package video

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"rest-gin/pkg/validators"
)

type handler struct {
	service VideoService
}

var validat *validator.Validate

func NewHandler(service VideoService) VideoHandler {
	validat = validator.New()
	validat.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &handler{
		service: service,
	}
}

func (h *handler) FindAll() []Video {
	return h.service.FindAll()
}
func (h *handler) Save(ctx *gin.Context) error {
	var video Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validat.Struct(video)
	if err != nil {
		return err
	}
	h.service.Save(video)
	return nil
}
