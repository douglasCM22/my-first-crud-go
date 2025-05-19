package service

import (
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/configuration/rest_err"
	model "my-first-crud-go/src/model/user"

	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(id string, user model.UserDomainInterface) *rest_err.RestErr {

	logs.Info("Updating user in the database", zap.String("Journey", "UpdateUser"))

	user.EncryptPassword()

	return nil
}
