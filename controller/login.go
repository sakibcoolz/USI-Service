package controller

import (
	"net/http"

	"USI-Service/model/restmodel"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Login(ctx *gin.Context) {
	var login restmodel.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		c.log.Error("Failed to bind json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.Login(login)
	if err != nil {
		c.log.Error("Failed to login", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
