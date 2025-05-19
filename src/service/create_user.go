package service

import (
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/configuration/rest_err"
	model "my-first-crud-go/src/model/user"

	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(user model.UserDomainInterface) *rest_err.RestErr {

	logs.Info("Creating user in the database", zap.String("Journey", "CreateUser"))
	user.EncryptPassword()

	return nil
}
