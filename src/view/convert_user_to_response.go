package view

import (
	"my-first-crud-go/src/controller/model/response"
	model "my-first-crud-go/src/model/user"
)

func ConvertUserToResponse(ud model.UserDomainInterface) *response.UserResponse {
	return &response.UserResponse{
		Id:       "test",
		Email:    ud.GetEmail(),
		Username: ud.GetUsername(),
		Age:      ud.GetAge(),
	}
}
