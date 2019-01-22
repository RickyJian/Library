package main

import (
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
	"library/controller"
	"net/http"
)

var b *controller.Book

func initServer() {
	router := gin.Default()
	t := template.TemplateWalk("assets/template/",".html")
	router.SetHTMLTemplate(t)
	router.StaticFS("/assets/static/",http.Dir("assets/static/"))
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.GET(projectPath+"/", index)
	router.Run()
}

func index (c *gin.Context){
	c.HTML(http.StatusOK, "index", gin.H{
	})
}
