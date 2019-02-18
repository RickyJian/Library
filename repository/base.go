package repository

import (
	"library/db"
)

type Base interface {
	Add() (isSuccessful bool)
	ReadByID()
	Update(data interface{})
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{})
	}
}

func Add(base Base) (isSuccessful bool) {
	return base.Add()
}

func ReadByID(base Base) {
	base.ReadByID()
}

func ReadAll(data interface{}) {
	db.GetConn().Find(data)
}

func Update(base Base , data interface{}){
	base.Update(data)
}