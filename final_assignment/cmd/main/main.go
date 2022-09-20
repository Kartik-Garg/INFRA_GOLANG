package main

import (
	"final/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := mux.NewRouter()
	// routes.RegisterBookStoreRoutes(r)
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:9010", r))

	router := gin.Default()

	routes.RegisterBookStoreRoutes(router)

	router.Run("localhost:9010")
}
