package main

import (
	"hanshal101/jwt/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetLoginHandler(t *testing.T) {
	mockResponse := `{"message":"using the GET Request"}`

	r := SetUpRouter()
	r.GET("/login", handlers.GetLoginHandler)

	req, _ := http.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSignupHandler(t *testing.T) {
	mockResponse := `{"message":"You have accessed the /signup GET path"}`

	r := SetUpRouter()
	r.GET("/signup", handlers.GetSignUpHandler)

	req, _ := http.NewRequest("GET", "/signup", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
