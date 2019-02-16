package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"library/model"
	"net/http"
	"strconv"
)

type Book struct {
	Name        string `json:"name"`
	Publication string `json:"publication"`
	Author      string `json:"author"`
	Price       string `json:"price"`
	Content     string `json:"content"`
}

var bookModel model.Book
var book Book

func (b *Book) New(c *gin.Context) {

	c.HTML(http.StatusOK, "book/new.html", gin.H{
		"path": "",
	})
}

func (b *Book) Add(c *gin.Context) {
	err := c.BindJSON(&book)
	copier.Copy(&bookModel,&book)
	price , err := strconv.Atoi(book.Price)
	bookModel.Price = int(price)
	if err == nil {
		bookModel.Add()
	}
}
