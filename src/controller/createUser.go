package controller

import (
	"fmt"
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/configuration/validation"
	"my-first-crud-go/src/controller/model/request"
	"my-first-crud-go/src/controller/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logs.Error("Error trying to bind json: ", err, zap.String("Journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		Id:       "1",
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Age:      userRequest.Age,
	}

	logs.Info("User created successfully", zap.String("Journey", "CreateUser"))

	c.JSON(http.StatusOK, response)
}
