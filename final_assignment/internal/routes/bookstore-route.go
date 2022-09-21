package routes

import (
	"final/internal/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine, b controllers.ControllerInterface) {
	router.POST("/book/", b.CreateBook)
	router.GET("/book/", b.GetAllBooks)
	router.GET("/book/:ID", b.GetBookById)
	router.DELETE("/book/:ID", b.DeleteBookById)
}
