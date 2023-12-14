package handlers

import (
	"context"
	"fmt"
	"hanshal101/jwt/internal/db"
	"hanshal101/jwt/internal/helpers"
	"hanshal101/jwt/models"
	jwttoken "hanshal101/jwt/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var collectionName = db.GetUsersCollection("users")

func SignUp(res *gin.Context) {
	var signUpObj *models.SignUp
	if err := res.BindJSON(&signUpObj); err != nil {
		fmt.Println("Error in Binding")
	}
	var claims = &models.JwtClaims{}
	claims.Email = signUpObj.Email
	claims.Phone = signUpObj.Phone
	claims.Username = signUpObj.Username
	claims.Audience = res.Request.Header.Get("Referer")
	var tokenCreation = time.Now().UTC()
	var expirationTime = tokenCreation.Add(time.Duration(10) * time.Hour)
	tokenString, err := jwttoken.GenerateToken(claims, expirationTime)

	if err != nil {
		fmt.Println("Error in generating token")
		return
	}

	hashedPass, err := helpers.EncryptPassword(signUpObj.Password)
	if err != nil {
		fmt.Println("error in hash pass")
	}

	signUpObj.Password = hashedPass

	result, err := collectionName.InsertOne(context.TODO(), signUpObj)
	if err != nil {
		fmt.Println("Error in Inserting Data")
	}
	fmt.Println(result)
	res.JSON(http.StatusCreated, tokenString)
}
