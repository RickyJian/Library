package controller

import (
	"github.com/gin-gonic/gin"
	"library/model"
	"net/http"
)

type Index struct {
}

var b model.Book

func (i *Index) Get(c *gin.Context) {
	bookSlice := b.ReadAll()
	c.HTML(http.StatusOK, "index", gin.H{
		"bookList": bookSlice,
	})
}
