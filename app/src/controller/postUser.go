package controller

import (
	"fmt"
	"golang-crud/app/src/configuration/app_errors"
	"golang-crud/app/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUserById(c *gin.Context) {

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := app_errors.NewBadRequestError("", nil)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
