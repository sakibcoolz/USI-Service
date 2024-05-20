package domain

import (
	"USI-Service/model/dbmodel"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IDomain interface {
	CreateUser(register *dbmodel.RegisterUser) error
	Login(username string, password string) (dbmodel.RegisterUser, error)
	UserDetails(ctx *gin.Context, username string) (dbmodel.RegisterUser, error)
}

type Domain struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDomain(db *gorm.DB, log *zap.Logger) IDomain {
	return &Domain{
		db:  db,
		log: log,
	}
}
