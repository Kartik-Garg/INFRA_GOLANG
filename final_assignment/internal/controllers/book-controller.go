package controllers

import (
	"final/internal/db"
	"final/internal/models"
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
	booksDB db.BooksDB
}

func NewBooksController(booksDB db.BooksDB) ControllerInterface {
	return &BooksDBimpl{
		booksDB,
	}
}

func (b *BooksDBimpl) CreateBook(c *gin.Context) {
	book := models.Book{}
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
	}

	res, err := b.booksDB.CreateBook(c, &book)
	if err != nil {
		c.JSON(400, gin.H{"error": "not able to create"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *BooksDBimpl) GetAllBooks(c *gin.Context) {
	books, err := b.booksDB.GetAllBooks(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b *BooksDBimpl) GetBookById(c *gin.Context) {
	id := c.Params.ByName("ID")
	book, err := b.booksDB.GetBookById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not get the book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *BooksDBimpl) DeleteBookById(c *gin.Context) {
	id := c.Params.ByName("ID")
	err := b.booksDB.DeleteBookById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
