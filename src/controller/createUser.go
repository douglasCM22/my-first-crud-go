package controller

import (
	"fmt"
	"my-first-crud-go/src/configuration/rest_err"
	"my-first-crud-go/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("invalid json body, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
