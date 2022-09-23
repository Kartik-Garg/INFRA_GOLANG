package main

import (
	"final/internal/controllers"
	"final/internal/routes"
	"final/internal/service"
	"final/pkg/db/mysql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//lets start with db, we can follow a bottom-up approach for easy understanding

	db, err := mysql.NewMYSQLDB()
	if err != nil {
		log.Fatal("Can not initialize the DB")
	}

	db.Migrate()

	booksSvc := service.NewBookService(db)

	booksController := controllers.NewBooksController(booksSvc)

	routes.RegisterBookStoreRoutes(router, booksController)

	router.Run("localhost:9010")

}
