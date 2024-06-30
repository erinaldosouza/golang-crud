package controller

import (
	"golang-crud/app/src/configuration/logger"
	"golang-crud/app/src/configuration/validation"
	"golang-crud/app/src/model"
	"golang-crud/app/src/model/request"
	"golang-crud/app/src/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	IUserDomainer model.IUserDomainer
)

func CreateUserById(c *gin.Context) {

	logger.Info("About to create user",
		zap.String("Journey", "CreateUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to create user", err)

		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email, userRequest.Pwd, userRequest.Name, userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   int(userRequest.Age),
	}

	logger.Info("User created succefully",
		zap.String("Journey", "CreateUser"),
	)

	c.JSON(http.StatusOK, response)
}
