package server

import (
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	t := template.TemplateWalk("assets/template/", ".html")
	router.SetHTMLTemplate(t)
	router.StaticFS("/assets/static/", http.Dir("assets/static/"))
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	controllerBind()
	router.Run()
}
