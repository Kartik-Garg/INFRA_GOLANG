package mysql

import (
	"context"
	"final/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//keep all the connections and implementation of all the methods here itself

// to avoid creating global variables, we can use structs
type MYSQLDB struct {
	*gorm.DB
}

// creating a new connection with mysql server
func NewMYSQLDB() (*MYSQLDB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &MYSQLDB{
		d,
	}, nil
}

// we can start adding all the db methods here
func (mysqlDB *MYSQLDB) Migrate() error {
	return mysqlDB.AutoMigrate(models.Book{})
}

func (mysqlDB *MYSQLDB) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	res := mysqlDB.Create(book)
	if res.Error != nil {
		return nil, res.Error
	}
	return book, nil
}

func (mysqlDB *MYSQLDB) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	books := []*models.Book{}
	res := mysqlDB.Find(&books)
	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func (mysqlDB *MYSQLDB) GetBookById(ctx context.Context, id string) (*models.Book, error) {
	book := &models.Book{}
	res := mysqlDB.Where("ID=?", id).Find(&book)

	if res.Error != nil {
		return nil, res.Error
	}
	return book, nil
}

func (mysqlDB *MYSQLDB) DeleteBookById(ctx context.Context, id string) error {
	book := &models.Book{}
	res := mysqlDB.Where("ID=?", id).Delete(&book)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
