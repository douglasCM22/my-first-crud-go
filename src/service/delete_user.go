package service

import (
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logs.Info("Deleting user in the database", zap.String("Journey", "DeleteUser"))
	return nil
}
