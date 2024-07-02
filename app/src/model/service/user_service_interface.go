package service

import (
	"golang-crud/app/src/configuration/app_errors"
	"golang-crud/app/src/model"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *app_errors.AppError
	UpdateUser(string, model.UserDomainInterface) *app_errors.AppError
	FindUser(string) (model.UserDomainInterface, *app_errors.AppError)
	DeleteUser(string) *app_errors.AppError
}
