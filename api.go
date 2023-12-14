package main

import (
	"fmt"
	"hanshal101/jwt/handlers"
	"hanshal101/jwt/internal/db"
	"hanshal101/jwt/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	db.MongoConnect()
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", handlers.GetLoginHandler)
	r.POST("/login", handlers.LoginHandler)

	r.GET("/api", middlewares.ValidateToken(), handlers.GetAPIHandler)

	r.GET("/signup", handlers.GetSignUpHandler)
	r.POST("/signup", handlers.SignUpHandler)

	r.Run(":9876")
}
