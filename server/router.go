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
	router.GET(s.Root+"/book/new/", b.New)
	router.POST(s.Root+"/book/add/", b.Add)
}
