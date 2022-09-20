package controllers

import (
	"final/pkg/models"

	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.JSON(200, newBooks)

}

func GetBookById(c *gin.Context) {
	id := c.Params.ByName("ID")
	//ID, err := strconv.ParseInt(id, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing string to int")
	// }
	bookDetails, _ := models.GetBookById(id)
	c.JSON(200, bookDetails)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.Bind(&book)
	b := book.CreateBook()
	c.JSON(200, b)

}
