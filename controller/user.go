package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"library/model"
	"library/util"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Account string
}

var userModel model.User

func (u *User) Add(c *gin.Context) {
	userModel.Name = c.PostForm("name")
	userModel.Account = c.PostForm("account")
	userModel.EMail = c.PostForm("email")
	gender, err := strconv.Atoi(c.PostForm("gender"))
	userModel.Gender = gender
	password := c.PostForm("password")
	validPassword := c.PostForm("validPassword")
	if strings.EqualFold(password, validPassword) && err == nil {
		userModel.Password = util.StringToSHA512(password)
		isSuccessful := userModel.Add()
		if isSuccessful {
			c.JSON(
				http.StatusOK,
				gin.H{
					"status":  200,
					"message": "success",
				},
			)
		}
	}
}

func (u *User) ReadByID(c *gin.Context)  {
	userModel.Account = c.PostForm("account")
	isRecordNotFound := userModel.ReadByID()
	copier.Copy(&b, book)
	if isRecordNotFound {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"message": "未新增使用者",
			},
		)
	}
}
