package config

import (
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
	"library/controller"
	"net/http"
)

type Server struct {
	Root string
}

var b *controller.Book
var i *controller.Index


func (s *Server)Init() {
	router := gin.Default()
	t := template.TemplateWalk("assets/template/", ".html")
	router.SetHTMLTemplate(t)
	router.StaticFS("/assets/static/", http.Dir("assets/static/"))
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.GET(s.Root+"/", i.Get)
	router.GET(s.Root+"/book/new/", b.New)
	router.Run()
}
