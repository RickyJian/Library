package repository

import (
	"library/db"
)

type Base interface {
	Add() (isSuccessful bool)
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{})
	}
}

func Add(base Base) (isSuccessful bool) {
	return base.Add()
}

func ReadAll(data interface{}){
	db.GetConn().Find(data)
}
