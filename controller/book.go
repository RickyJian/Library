package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct{

}

func (b *Book)Index(c *gin.Context) {
	c.HTML(http.StatusOK,"book/index.html",gin.H{
		"title":"首頁",
	})
}