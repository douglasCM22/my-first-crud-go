package controller

import (
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/configuration/validation"
	"my-first-crud-go/src/controller/model/request"
	model "my-first-crud-go/src/model/user"
	"my-first-crud-go/src/view"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logs.Error("Error trying to bind json: ", err, zap.String("Journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Username, userRequest.Email, userRequest.Password, userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {

		logs.Error("Fail to create a user", err, zap.String("Journey", "CreateUser"))

		c.JSON(err.Code, err)
		return
	}

	logs.Info("User created successfully", zap.String("Journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertUserToResponse(domain))

}
