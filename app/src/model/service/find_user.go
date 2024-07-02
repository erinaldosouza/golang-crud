package service

import (
	"golang-crud/app/src/configuration/app_errors"
)

// FindUser implements IUserDomain.
func (ud *userDomain) FindUser(string) (*userDomain, *app_errors.AppError) {
	panic("unimplemented")
}
