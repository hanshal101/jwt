package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hanshal101/jwt/internal/helpers"
	"github.com/hanshal101/jwt/models"
	jwttoken "github.com/hanshal101/jwt/token"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// var collectionName = db.GetUsersCollection("users")

func LoginHandler(res *gin.Context) {
	var loginObj models.LoginRequest
	if err := res.BindJSON(&loginObj); err != nil {
		fmt.Println("Error in Binding ", err)
	}

	var user models.VerifyLogin
	filter := bson.D{{Key: "username", Value: &loginObj.Username}}
	err := collectionName.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"message": "Invalid username or password"})
		return
	}

	valid := helpers.ValidatePassword(user.Password, loginObj.Password)

	if !valid {
		res.JSON(http.StatusBadRequest, gin.H{"message": "password incorrect"})
		return
	}
	// fmt.Println(result)
	var claims = &models.JwtClaims{}
	claims.Email = user.Email
	claims.Phone = user.Phone
	claims.Audience = res.Request.Header.Get("Referer")
	claims.Username = loginObj.Username

	var tokenCreation = time.Now().UTC()
	var expirationTime = tokenCreation.Add(time.Duration(30) * time.Second)
	tokenString, err := jwttoken.GenerateToken(claims, expirationTime)

	if err != nil {
		fmt.Println("Error in generating token")
		return
	}
	res.SetCookie("apikey", tokenString, int(expirationTime.Unix()), "/", "/", true, true)

	res.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetLoginHandler(res *gin.Context) {
	res.JSON(http.StatusOK, gin.H{"message": "using the GET Request"})
}
