package model

import (
	"fmt"
	"github.com/jinzhu/copier"
	"library/repository"
)

type User struct {
	Name     string
	Account  string
	Password string
	Gender   int
	EMail    string
}

const roleID = 0

func (u *User) Add() (isSuccessful bool) {
	user := &repository.User{}
	copier.Copy(user, &u)
	user.Role = roleID
	isSuccessful = repository.Add(user)
	fmt.Println(isSuccessful)
	return
}

func (u *User) ReadByID() (isRecordNotFound bool) {
	user := &repository.User{Account: u.Account}
	isRecordNotFound = repository.ReadByID(user)
	copier.Copy(&u, user)
	return
}

func (u *User) ReadForLogin() (isRecordNotFound bool) {
	user := &repository.User{Account: u.Account,Password:u.Password}
	isRecordNotFound = user.ReadForLogin()
	copier.Copy(&u, user)
	return
}

