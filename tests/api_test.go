package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hanshal101/jwt/handlers"

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

// func TestSignUpHandler(t *testing.T) {

// 	r := SetUpRouter()
// 	r.POST("/signup", handlers.SignUpHandler)
// 	user := models.SignUp{
// 		Username: "test",
// 		Password: "654321",
// 		Email:    "rohan@mail.com",
// 		Phone:    "879546213",
// 	}
// 	jsonValue, _ := json.Marshal(&user)
// 	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// }

// func TestLoginHandler(t *testing.T) {
// 	r := SetUpRouter()
// 	r.POST("/login", handlers.LoginHandler)
// 	loginUser := models.LoginRequest{
// 		Username: "test",
// 		Password: "654321",
// 	}
// 	jsonValue, err := json.Marshal(&loginUser)
// 	assert.NoError(t, err)

// 	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
// 	assert.NoError(t, err)

// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response map[string]string
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// }
