package routes

import (
	"final/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:ID", controllers.GetBookById)
}
