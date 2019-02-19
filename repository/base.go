package repository

import (
	"library/db"
)

type Base interface {
	Add() (isSuccessful bool)
	ReadByID() (isRecordNotFound bool)
	Update(data interface{})
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{},&User{})
	}
}

func Add(base Base) (isSuccessful bool) {
	return base.Add()
}

func ReadByID(base Base) (isRecordNotFound bool) {
	return base.ReadByID()
}

func ReadAll(data interface{}) {
	db.GetConn().Find(data)
}

func Update(base Base, data interface{}) {
	base.Update(data)
}
