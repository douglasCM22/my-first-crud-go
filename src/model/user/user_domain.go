package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	email    string
	password string
	username string
	age      int8
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetUsername() string
	GetAge() int8
	EncryptPassword()
}

func NewUserDomain(email, password, username string, age int8) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
		username: username,
		age:      age,
	}
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetUsername() string {
	return ud.username
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hashEncode := md5.New()
	defer hashEncode.Reset()
	hashEncode.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hashEncode.Sum(nil))
}
