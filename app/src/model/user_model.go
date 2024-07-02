package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPwd() string
	GetName() string
	GetAge() int8
}

func NewUserDomain(
	email, pwd, name string,
	age int8,
) IUserDomainer {
	return &userDomain{
		email, pwd, name, age,
	}
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPwd() string {
	return ud.pwd
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

type userDomain struct {
	email string
	pwd   string
	name  string
	age   int8
}

func (ud *userDomain) encryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.pwd))
	ud.pwd = hex.EncodeToString(hash.Sum(nil))
}
