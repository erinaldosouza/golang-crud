package model

import (
	"crypto/md5"
	"encoding/hex"
	"golang-crud/app/src/configuration/app_errors"
)

func NewUserDomain(
	email, pwd, name string,
	age int8,
) IUserDomainer {
	return &userDomain{
		email, pwd, name, age,
	}
}

type userDomain struct {
	Email string
	Pwd   string
	Name  string
	Age   int8
}

type IUserDomainer interface {
	CreateUser() *app_errors.AppError
	UpdateUser(string) *app_errors.AppError
	FindUser(string) (*userDomain, *app_errors.AppError)
	DeleteUser(string) *app_errors.AppError
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Pwd))
	ud.Pwd = hex.EncodeToString(hash.Sum(nil))
}
