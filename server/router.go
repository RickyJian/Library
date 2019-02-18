package server

import (
	"library/config"
	"library/controller"
)

var b *controller.Book
var i *controller.Index

func controllerBind() {
	s := config.GetProject()
	router.GET(s.Root+"/", i.Get)
	router.GET(s.Root+"/book/detail/:id", b.Detal)
	router.GET(s.Root+"/book/new/", b.New)
	router.GET(s.Root+"/book/edit/:id/", b.Edit)
	router.POST(s.Root+"/book/add/", b.Add)
}
