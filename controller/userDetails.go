package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UserDetails(ctx *gin.Context) {
	user := ctx.Param("user")

	register := c.service.UserDetails(ctx, user)

	ctx.JSON(http.StatusOK, register)
}
