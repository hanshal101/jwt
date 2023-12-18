package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwttoken "github.com/hanshal101/jwt/token"

	"github.com/hanshal101/jwt/internal/db"
	"github.com/hanshal101/jwt/internal/helpers"
	"github.com/hanshal101/jwt/models"

	"github.com/gin-gonic/gin"
)

var collectionName = db.GetUsersCollection("users")

func SignUpHandler(res *gin.Context) {

	var signUpObj *models.SignUp
	if err := res.BindJSON(&signUpObj); err != nil {
		fmt.Println("Error in Binding")
	}
	isEmailPhoneValid, err := helpers.ValidEmailPhone(signUpObj.Email, signUpObj.Phone)
	if err != nil {
		fmt.Printf("Error checking email and phone validity: %v\n", err)
		res.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	if !isEmailPhoneValid {
		res.JSON(http.StatusBadRequest, gin.H{"message": "Email or Phone already exist"})
		return
	}
	var claims = &models.JwtClaims{}
	claims.Email = signUpObj.Email
	claims.Phone = signUpObj.Phone
	claims.Username = signUpObj.Username
	var tokenCreation = time.Now().UTC()
	var expirationTime = tokenCreation.Add(time.Duration(30) * time.Second)
	claims.Audience = "192.168.0.107"
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

func GetSignUpHandler(res *gin.Context) {
	res.JSON(http.StatusOK, gin.H{"message": "You have accessed the /signup GET path"})
}
