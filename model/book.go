package model

import (
	"github.com/jinzhu/copier"
	"library/repository"
)

type Book struct {
	ID          int
	Name        string
	Publication string
	Author      string
	Price       int
	Content     string
	Cover       string
}

func (b *Book) Add() (isSuccessful bool) {
	book := &repository.Book{}
	copier.Copy(&book, &b)
	isSuccessful = repository.Add(book)
	if !isSuccessful {
		return isSuccessful
	}
	copier.Copy(&b, &book)
	return isSuccessful
}

func (b *Book) ReadAll() (resultSlice []Book) {
	var bookSlice = []repository.Book{}
	repository.ReadAll(&bookSlice)
	copier.Copy(&resultSlice, &bookSlice)
	return
}
func (b *Book) ReadByID() {
	book := &repository.Book{ID: b.ID}
	repository.ReadByID(book)
	copier.Copy(&b, book)
}
