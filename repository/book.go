package repository

import (
	"library/db"
	"time"
)

type Book struct {
	ID          int       `gorm:"type:serial;primary_key"`
	Name        string    `gorm:"not null"`
	Publication string    `gorm:"not null"`
	Author      string    `gorm:"not null"`
	Price       int       `gorm:"not null"`
	Content     string    `gorm:"null"`
	Image       string    `gorm:"null"`
	CreateTime  time.Time `gorm:"not null"`
}

func (b *Book) Add() bool {
	if db.GetConn().NewRecord(b) {
		db.GetConn().Create(&b)
		return true
	}
	return false
}
