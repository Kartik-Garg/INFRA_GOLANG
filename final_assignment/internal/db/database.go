package db

import (
	"context"
	"final/internal/models"
)

// here we can write interface for all the db related methods
type BooksDB interface {
	GetBookById(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) (*[]models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	DeleteBookById(ctx context.Context, id string) error
}
