package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct{

}

func (b *Book)New(c *gin.Context) {
	c.HTML(http.StatusOK,"book/new.html",gin.H{
		"title":"首頁",
	})
}