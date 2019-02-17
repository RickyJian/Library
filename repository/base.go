package repository

import "library/db"

type Base interface {
	Add() (isSuccessful bool)
	ReadAll() (dataList []interface{})
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{})
	}
}

func Add(base Base) (isSuccessful bool) {
	return base.Add()
}

func ReadAll(base Base) (dataList []interface{}) {
	return base.ReadAll()
}
