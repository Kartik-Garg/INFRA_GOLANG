package main

import (
	"final/models"
	"final/routes"
)

func main() {
	models.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run()
}
