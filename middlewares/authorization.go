package middlewares

import (
	"fmt"
	"hanshal101/jwt/models"
	"hanshal101/jwt/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnUnauth(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error:   fmt.Errorf("Unauth"),
		Status:  http.StatusUnauthorized,
		Message: "You are unauthorized",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		referer := context.Request.Header.Get("Referer")

		valid, claims := token.VerifyToken(tokenString, referer)

		if !valid {
			ReturnUnauth(context)
		}

		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["Username"] = claims.Username
	}
}
