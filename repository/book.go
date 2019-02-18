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
	Cover       string    `gorm:"not null;default:'/assets/static/image/dbook.jpg'"`
	CreateTime  time.Time `gorm:"not null"`
}

func (b *Book) Add() (isSuccessful bool) {
	if db.GetConn().NewRecord(b) {
		b.CreateTime = time.Now()
		db.GetConn().Create(&b)
		return true
	}
	return false
}

func (b *Book) ReadByID() {
	db.GetConn().Where("ID = ?", b.ID).First(&b)
}

func (b *Book) Update(data interface{}){
	db.GetConn().Model(&b).Where("ID = ?", b.ID).UpdateColumns(data)
}

