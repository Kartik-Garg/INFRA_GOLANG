package service

import (
	"context"
	"final/internal/db"
	"final/internal/models"
)

type BookService interface {
	GetBookById(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) (*[]models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	DeleteBookById(ctx context.Context, id string) error
}

type BookServiceImpl struct {
	booksDB db.BooksDB
}

func NewBookService(booksDB db.BooksDB) BookService {
	return &BookServiceImpl{
		booksDB,
	}
}

func (bs *BookServiceImpl) GetBookById(ctx context.Context, id string) (*models.Book, error) {
	return bs.booksDB.GetBookById(ctx, id)
}

func (bs *BookServiceImpl) GetAllBooks(ctx context.Context) (*[]models.Book, error) {
	return bs.booksDB.GetAllBooks(ctx)
}

func (bs *BookServiceImpl) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	return bs.booksDB.CreateBook(ctx, book)
}

func (bs *BookServiceImpl) DeleteBookById(ctx context.Context, id string) error {
	return bs.booksDB.DeleteBookById(ctx, id)
}
