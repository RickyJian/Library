package controller

import (
	"github.com/gin-gonic/gin"
	"library/model"
	"net/http"
)

type System struct {
}

var b model.Book

func (s *System) Index(c *gin.Context) {
	bookSlice := b.ReadAll()
	c.HTML(http.StatusOK, "index", gin.H{
		"bookList": bookSlice,
	})
}

func (s *System) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"path": "",
	})
}

func (s *System) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", gin.H{
		"path": "",
	})
}
