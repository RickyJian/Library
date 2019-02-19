package repository

import (
	"library/db"
)

type User struct {
	Account  string `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Gender   int    `gorm:"DEFAULT:'1';not null"`
	EMail    string `gorm:"null"`
	Role     int    `gorm:"DEFAULT:'0';not null"`
}

func (u *User) Add() (isSuccessful bool) {
	if db.GetConn().Where("Account = ? ", u.Account).Find(&u).RecordNotFound() {
		db.GetConn().Create(&u)
		return true
	}
	return false
}

func (u *User) ReadByID() (isRecordNotFound bool) {
	isRecordNotFound = db.GetConn().Where("Account = ?", u.Account).Find(&u).RecordNotFound()
	return
}

func (u *User) Update(data interface{}) {
	db.GetConn().Model(&u).Where("Account = ?", u.Account).UpdateColumns(data)
}
