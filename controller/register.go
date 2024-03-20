package controller

import (
	"USI-Service/model/restmodel"
	"USI-Service/utils/zerrors"
	"USI-Service/utils/zerrors/apperrors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Register(ctx *gin.Context) {
	var register *restmodel.RegisterUser

	if err := ctx.ShouldBindJSON(&register); err != nil {
		c.log.Error("Failed to bind JSON", zap.Error(err))

		err := zerrors.Errors(apperrors.BadRequest, err)

		err.(*zerrors.AppError).Response(ctx)

		return
	}

	if err := c.validator.Struct(register); err != nil {
		c.log.Error("Failed to validate JSON", zap.Error(err))

		err := zerrors.Errors(apperrors.BadRequest, err)

		err.(*zerrors.AppError).Response(ctx)

		return
	}

	resp, err := c.service.RegisterUser(*register)
	if err != nil {
		c.log.Error("Failed to register user", zap.Error(err))

		err.(*zerrors.AppError).Response(ctx)

		return
	}

	ctx.JSON(201, resp)
}
