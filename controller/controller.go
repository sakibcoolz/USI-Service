package controller

import (
	"USI-Service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type IController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UserDetails(ctx *gin.Context)
}

type Controller struct {
	service   service.IService
	log       *zap.Logger
	validator *validator.Validate
}

func NewController(service service.IService, log *zap.Logger, validator *validator.Validate) IController {
	return &Controller{
		service,
		log,
		validator,
	}
}
