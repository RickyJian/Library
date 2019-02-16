package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	form, _ := c.MultipartForm()
	fmt.Println(form.File["images"])
	bookModel.Name = form.Value["name"][0]
	bookModel.Author = form.Value["author"][0]
	bookModel.Content = form.Value["content"][0]
	bookModel.Publication = form.Value["publication:"][0]
	price, err := strconv.Atoi(form.Value["price"][0])
	bookModel.Price = price
	if err == nil {
		bookModel.Add()
	}
}
