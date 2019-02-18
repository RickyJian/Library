package controller

import (
	"github.com/gin-gonic/gin"
	"library/model"
	"library/util"
	"net/http"
	"strconv"
)

type Book struct {
	Name        string `json:"name"`
	Publication string `json:"publication"`
	Author      string `json:"author"`
	Price       string `json:"price"`
	Content     string `json:"content"`
}

var bookModel model.Book
var book Book

func (b *Book) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		bookModel.ID = id
		isRecordNotFound := bookModel.ReadByID()
		if !isRecordNotFound {
			c.HTML(http.StatusOK, "book/edit.html", gin.H{
				"book": bookModel,
			})
		}
	}
	c.HTML(http.StatusNotFound, "404.html", gin.H{
	})

}
func (b *Book) New(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new.html", gin.H{
		"path": "",
	})
}

func (b *Book) Add(c *gin.Context) {
	form, _ := c.MultipartForm()
	bookModel.Name = form.Value["name"][0]
	bookModel.Author = form.Value["author"][0]
	bookModel.Content = form.Value["content"][0]
	bookModel.Publication = form.Value["publication"][0]
	price, err := strconv.Atoi(form.Value["price"][0])
	bookModel.Price = price
	bookModel.Cover = "C:\\tmp\\upload\\" + form.File["images"][0].Filename
	if err == nil {
		isSuccessful := bookModel.Add()
		if isSuccessful {
			err = util.CreateFile(strconv.Itoa(bookModel.ID), *form.File["images"][0])
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

func (b *Book) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		bookModel.ID = id
		bookModel.Content = c.PostForm("content")
		bookModel.Name = c.PostForm("name")
		bookModel.Author = c.PostForm("author")
		price, _ := strconv.Atoi(c.PostForm("price"))
		bookModel.Price = price
		bookModel.Publication = c.PostForm("publication")
		bookModel.Update()
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"message": "success",
			},
		)
	}
}

func (b *Book) UpdateImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		bookModel.ID = id
		bookModel.Cover = "C:\\tmp\\upload\\" + form.File["images"][0].Filename
		bookModel.Update()
		err = util.CreateFile(strconv.Itoa(bookModel.ID), *form.File["images"][0])
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"message": "success",
			},
		)
	}
}

func (b *Book) Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		bookModel.ID = id
		isRecordNotFound := bookModel.ReadByID()
		if !isRecordNotFound {
			c.HTML(http.StatusOK, "book/detail.html", gin.H{
				"book": bookModel,
			})
		}
	}
	c.HTML(http.StatusNotFound, "404.html", gin.H{
	})
}
