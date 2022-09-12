package services

import "github.com/gin-gonic/gin"

type BookService interface {
	CreateBook(*gin.Context)
	GetAll(*gin.Context)
}
