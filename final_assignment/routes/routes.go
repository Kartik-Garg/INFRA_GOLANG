package routes

import (
	"final/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", services.FindBooks)
	router.POST("/books", services.CreateBook)
	return router
}
