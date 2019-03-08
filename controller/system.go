package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/model"
	"library/util"
	"net/http"
)

type System struct {
}

var b model.Book
var u model.User

func (s *System) Index(c *gin.Context) {
	bookSlice := b.ReadAll()
	c.HTML(http.StatusOK, "index", gin.H{
		"bookList": bookSlice,
	})
}

func (s *System) Login(c *gin.Context) {
	u.Account = c.PostForm("account")
	u.Password = util.StringToSHA512(c.PostForm("password"))
	fmt.Printf("%s Account \n",u.Account )
	fmt.Printf("%s Password",u.Password )
	isRecordNotFound := u.ReadForLogin()
	if !isRecordNotFound{
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"isSuccess": true ,
				"message": "登入成功",
			},
		)
	}else{
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"isSuccess": false ,
				"message": "登入失敗",
			},
		)
	}
}

func (s *System) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin", gin.H{
		"path": "",
	})
}

func (s *System) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", gin.H{
		"path": "",
	})
}
