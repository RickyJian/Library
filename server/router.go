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
	router.GET(s.Root+"/book/:action/", b.Dispatch)
}
