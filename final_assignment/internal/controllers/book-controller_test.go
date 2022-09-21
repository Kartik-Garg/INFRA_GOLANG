package controllers

import (
	"bytes"
	"encoding/json"
	"final/internal/models"
	"final/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateBook(t *testing.T) {
	booksSvcMock := mocks.NewBooksService(t)

	//need controller as well, and we will pass in the mocked service to the controller
	booksCtrl := NewBooksController(booksSvcMock)

	book := models.Book{
		Name:        "mockname1",
		Author:      "mockauthor1",
		Publication: "mockpublication1",
	}

	t.Run("200 and valid book in return", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		c.Request = NewCreateBookRequest(&book)
		require.NotNil(t, c.Request)

		createdBookCalled := 0

		//now we need to run the createbook method which is in our mock service
		booksSvcMock.On("CreateBook", mock.Anything, mock.Anything).Return(&book, nil).Run(func(args mock.Arguments) {
			createdBookCalled++
		}).Once()
		//now we have run the mock for the service, but the request goes from controller and then to service, so we run the controller now
		booksCtrl.CreateBook(c)

		require.Equal(t, http.StatusOK, c.Writer.Status())
	})
}

func TestGetAllBooks(t *testing.T) {
	//we need mock service, so first we have to register that with the mock service layer and then a controller
	booksSvcMock := mocks.NewBooksService(t)
	//need a new controller now
	booksCtrl := NewBooksController(booksSvcMock)

	t.Run("should return 200 and all the books", func(t *testing.T) {
		//we need a request first on which we can run the controller and the service
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = NewGetAllBooksRequest()
		getAllBooksCalled := 0
		//start the mock service
		booksSvcMock.On("GetAllBooks", mock.Anything, mock.Anything).Return([]*models.Book{}, nil).Run(func(args mock.Arguments) {
			getAllBooksCalled++
		}).Once()

		//now since service is set we need to pass the request to the controller
		booksCtrl.GetAllBooks(c)
		require.Equal(t, 1, getAllBooksCalled)
		require.Equal(t, http.StatusOK, c.Writer.Status())
	})
}

func NewCreateBookRequest(book *models.Book) *http.Request {
	data, err := json.Marshal(book)
	if err != nil {
		return nil
	}
	req, _ := http.NewRequest(http.MethodPost, "/book/", bytes.NewBuffer(data))
	return req
}

func NewGetAllBooksRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/book/", nil)
	return req
}
