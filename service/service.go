package service

import (
	"USI-Service/domain"
	"USI-Service/model/restmodel"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IService interface {
	RegisterUser(register restmodel.RegisterUser) (response restmodel.RegisterUserResponse, err error)
	Login(login restmodel.Login) (response restmodel.LoginUserResponse, err error)
	UserDetails(ctx *gin.Context, user string) restmodel.RegisterUser
}

type Service struct {
	store domain.IDomain
	log   *zap.Logger
}

func NewService(store domain.IDomain, log *zap.Logger) IService {
	return &Service{
		store: store,
		log:   log,
	}
}
