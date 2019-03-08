package server

import (
	"library/config"
	"library/controller"
)

var b *controller.Book
var sys *controller.System
var u *controller.User

func controllerBind() {
	s := config.GetProject()
	// 系統
	router.GET(s.Root+"/", sys.Index)
	router.GET(s.Root+"/signup/", sys.SignUp)
	router.GET(s.Root+"/signin/", sys.SignIn)
	router.POST(s.Root+"/login/", sys.Login)
	// 使用者
	router.POST(s.Root+"/user/read/:account/", u.ReadByID)
	router.POST(s.Root+"/user/add/", u.Add)
	// 書籍系統
	router.GET(s.Root+"/book/detail/:id/", b.Detail)
	router.GET(s.Root+"/book/new/", b.New)
	router.GET(s.Root+"/book/edit/:id/", b.Edit)
	router.POST(s.Root+"/book/add/", b.Add)
	router.POST(s.Root+"/book/detail/:id/image/upload/", b.UpdateImage)
	router.PUT(s.Root+"/book/detail/:id/", b.Update)
}
