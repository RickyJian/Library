package controller

import (
	"github.com/gin-gonic/gin"
	"library/model"
	"net/http"
	"strconv"
)

type User struct {
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
	if password == validPassword && err == nil {
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
