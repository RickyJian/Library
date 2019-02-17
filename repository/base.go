package repository

import "library/db"

type Base interface {
	Add() (data interface{}, success bool)
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{})
	}
}

func Add(base Base) (data interface{}, isSuccessful bool) {
	return base.Add()
}
