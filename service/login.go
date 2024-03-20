package service

import (
	"USI-Service/model/dbmodel"
	"USI-Service/model/restmodel"

	"go.uber.org/zap"
)

func (s *Service) Login(login restmodel.Login) (response restmodel.LoginUserResponse, err error) {
	s.log.Info("Logging in user")
	var registered dbmodel.RegisterUser
	registered, err = s.store.Login(login.Username, login.Password)
	if err != nil {
		s.log.Error("Failed to login", zap.Error(err))

		return response, err
	}

	response = restmodel.LoginUserResponse{
		LoggedIn: true,
		Success:  "User logged in successfully",
		Data: restmodel.Data{
			Name:  registered.Name,
			Email: registered.Email,
			Token: "token",
		},
	}

	return response, nil
}
