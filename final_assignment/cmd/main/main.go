package main

import (
	"final/internal/controllers"
	"final/internal/db/mysql"
	"final/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes.RegisterBookStoreRoutes(router)

	// router.Run("localhost:9010")

	//lets start with db

	db, err := mysql.NewMYSQLDB()
	if err != nil {
		log.Fatal("Can not initialize the DB")
	}

	db.Migrate()

	booksController := controllers.NewBooksController(db)

	routes.RegisterBookStoreRoutes(router, booksController)

	router.Run("localhost:9010")

}
