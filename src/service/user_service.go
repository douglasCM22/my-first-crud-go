package service

import (
	"my-first-crud-go/src/configuration/rest_err"
	model "my-first-crud-go/src/model/user"
)

type userDomainService struct{}

func NewUserDomainService() UserServiceInterface {
	return &userDomainService{}
}

type UserServiceInterface interface {
	CreateUser(user model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(id string, user model.UserDomainInterface) *rest_err.RestErr
	FindUser(val string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
