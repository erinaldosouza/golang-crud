package service

import (
	"golang-crud/app/src/configuration/app_errors"
	"golang-crud/app/src/configuration/logger"

	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *app_errors.AppError {
	logger.Info("Init createUser model", zap.String("journey", "createUSer"))
	ud.EncryptPassword()
	logger.Info("Finished createUser model", zap.String("journey", "createUser"))
	return nil
}
