package main

import (
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
	"library/controller"
	"net/http"
)

var b *controller.Book
var i *controller.Index

func initServer() {
	router := gin.Default()
	t := template.TemplateWalk("assets/template/",".html")
	router.SetHTMLTemplate(t)
	router.StaticFS("/assets/static/",http.Dir("assets/static/"))
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.GET(projectPath+"/", i.Get)
	router.GET(projectPath+"/book/new/", b.New)
	router.Run()
}

