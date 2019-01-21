package main

import (
	"fmt"
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
	"library/controller"
	"net/http"
)

var b *controller.Book

func initServer() {
	router := gin.Default()
	router.Static("/assets/","assets/static/")
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	t := template.TemplateWalk("assets/template/",".html")
	router.SetHTMLTemplate(t)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("test")
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "首頁",
		})
	})
	router.GET("/inner", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inner/index", gin.H{
			"title": "首頁",
		})
	})

	router.GET("/book/", b.Index)
	router.Run()
}
