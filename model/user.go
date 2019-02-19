package model

type User struct {
	Name     string
	Account  string
	Password string
	Gender   int
	EMail    string
}

func (u *User) Add() (isSuccessful bool) {
	return
}
