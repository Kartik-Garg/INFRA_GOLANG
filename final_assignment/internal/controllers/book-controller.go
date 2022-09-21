package controllers

import (
	"final/internal/models"
	"final/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	CreateBook(c *gin.Context)
	GetAllBooks(c *gin.Context)
	GetBookById(c *gin.Context)
	DeleteBookById(c *gin.Context)
}

type BooksDBimpl struct {
	booksSvc service.BookService
}

func NewBooksController(booksSvc service.BookService) ControllerInterface {
	return &BooksDBimpl{
		booksSvc,
	}
}

func (b *BooksDBimpl) CreateBook(c *gin.Context) {
	book := models.Book{}
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
	}

	res, err := b.booksSvc.CreateBook(c, &book)
	if err != nil {
		c.JSON(400, gin.H{"error": "not able to create"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *BooksDBimpl) GetAllBooks(c *gin.Context) {
	books, err := b.booksSvc.GetAllBooks(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b *BooksDBimpl) GetBookById(c *gin.Context) {
	id := c.Params.ByName("ID")
	book, err := b.booksSvc.GetBookById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not get the book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *BooksDBimpl) DeleteBookById(c *gin.Context) {
	id := c.Params.ByName("ID")
	err := b.booksSvc.DeleteBookById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
