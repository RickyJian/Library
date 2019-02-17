package model

import (
	"fmt"
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
	fmt.Println(resultSlice[0].Cover)
	return
}
