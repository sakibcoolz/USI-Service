package service

import (
	"USI-Service/model/restmodel"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) UserDetails(ctx *gin.Context, user string) restmodel.RegisterUser {
	register, err := s.store.UserDetails(ctx, user)
	if err != nil {
		s.log.Error("failed to get error", zap.Error(err))

		return restmodel.RegisterUser{}
	}

	return restmodel.RegisterUser{
		Username: register.Email,
		Email:    register.Email,
		Password: "",
		Name:     register.Name,
		Role:     register.Role,
	}

}
