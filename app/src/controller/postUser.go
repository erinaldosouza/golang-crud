package controller

import (
	"fmt"
	"golang-crud/app/src/configuration/validation"
	"golang-crud/app/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUserById(c *gin.Context) {

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Println(err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
