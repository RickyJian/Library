package repository

import "library/db"

type Base interface {
	Add() bool
}

func init() {
	if db.IsMigrate {
		db.GetConn().AutoMigrate(&Book{})
	}
}

func Add(base Base) bool {
	return base.Add()
}
