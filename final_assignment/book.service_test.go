package main

import (
	"bytes"
	"final/models"
	"final/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {

	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestFindBooks(t *testing.T) {
	models.ConnectDatabase()

	request, _ := http.NewRequest("GET", "/books", nil)
	response := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(response, request)

	expected := `{"data":[{"id":1,"title":"Start with why","author":"Simon Sinek"},{"id":2,"title":"GOW","author":"PS"},{"id":3,"title":"GOW","author":"PS"}]}`

	assert.Equal(t, expected, response.Body.String())
}

func TestCreateBook(t *testing.T) {
	models.ConnectDatabase()

	// book := models.Book{
	// 	4,
	// 	"Alchemist",
	// 	"Paulo",
	// }

	var jsonStr = []byte(`{"id":4,"title":"Alchemist","author":"paulo"}`)

	request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(response, request)

	//passing data

	assert.Equal(t, http.StatusOK, response.Code)

}
