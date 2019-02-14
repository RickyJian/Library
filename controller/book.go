package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name string `json:"name"`
	Publication string `json:"publication"`
	Author string `json:"author"`
	Price string `json:"price"`
	Content string `json:"content"`
}

func (b *Book) New(c *gin.Context) {

	c.HTML(http.StatusOK, "book/new.html", gin.H{
		"path":"",
	})
}

func (b *Book)Add(c *gin.Context){
	var book Book;
	err := c.BindJSON(&book)
	if err != nil {

	}
}
