package domain

import (
	"USI-Service/model/dbmodel"
	"USI-Service/utils/zerrors"
	"USI-Service/utils/zerrors/apperrors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (d *Domain) CreateUser(register *dbmodel.RegisterUser) error {
	if err := d.db.Create(&register).Error; err != nil {
		d.log.Error("Failed to create user", zap.Error(err))

		return zerrors.Errors(apperrors.DatabaseError, err)
	}

	return nil
}

func (d *Domain) Login(username string, password string) (dbmodel.RegisterUser, error) {
	var user dbmodel.RegisterUser
	if err := d.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		d.log.Error("Failed to login", zap.Error(err))

		return user, zerrors.Errors(apperrors.DatabaseError, err)
	}

	return user, nil
}

func (d *Domain) UserDetails(ctx *gin.Context, username string) (dbmodel.RegisterUser, error) {
	var register dbmodel.RegisterUser

	if err := d.db.Where("username = ?", username).First(&register).Error; err != nil {
		d.log.Error("Failed to login", zap.Error(err))

		return register, zerrors.Errors(apperrors.DatabaseError, err)
	}
	return register, nil
}
