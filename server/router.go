package server

import (
	"library/config"
	"library/controller"
)

var b *controller.Book
var sys *controller.System

func controllerBind() {
	s := config.GetProject()
	router.GET(s.Root+"/", sys.Index)
	router.GET(s.Root+"/signup/", sys.SignUp)
	router.GET(s.Root+"/login/", sys.Login)
	router.GET(s.Root+"/book/detail/:id/", b.Detail)
	router.GET(s.Root+"/book/new/", b.New)
	router.GET(s.Root+"/book/edit/:id/", b.Edit)
	router.POST(s.Root+"/book/add/", b.Add)
	router.PUT(s.Root+"/book/detail/:id/", b.Update)
	router.POST(s.Root+"/book/detail/:id/image/upload/", b.UpdateImage)
}
