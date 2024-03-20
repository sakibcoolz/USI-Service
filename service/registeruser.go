package service

import (
	"USI-Service/model/dbmodel"
	"USI-Service/model/restmodel"

	"go.uber.org/zap"
)

func (s *Service) RegisterUser(register restmodel.RegisterUser) (response restmodel.RegisterUserResponse, err error) {
	var registerDB dbmodel.RegisterUser

	// Copy the struct
	registerDB.Name = register.Name
	registerDB.Username = register.Username
	registerDB.Password = register.Password
	registerDB.Email = register.Email
	registerDB.Role = register.Role
	s.log.Info("Registering user")
	s.log.Info("Registering user", zap.Any("username", registerDB))

	err = s.store.CreateUser(&registerDB)
	if err != nil {
		s.log.Error("Failed to create user", zap.Error(err))

		return response, err
	}

	response.Username = register.Username
	response.Stored = true

	return response, nil
}
