package middlewares

import (
	"net/http"
	"github.com/hanshal101/jwt/models"
	"github.com/hanshal101/jwt/token"
	"github.com/gin-gonic/gin"
)

type UnauthorizedError string

func (e UnauthorizedError) Error() string {
	return string(e)
}

type Response struct {
	Error   error  `json:"error,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ReturnUnauth(context *gin.Context) {
	unauthorizedError := UnauthorizedError("UnauthorizedError")

	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error:   unauthorizedError,
		Status:  http.StatusUnauthorized,
		Message: "You are unauthorized",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Request.Cookie("apikey")
		if err != nil {
			ReturnUnauth(context)
		}
		if len(cookie.Value) == 0 {
			// Handle the case where the cookie is not found
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token not found",
				"message": "You are unauthorized",
			})
			return
		}
		tokenString := cookie.Value
		referer := "192.168.0.107"

		valid, claims := token.VerifyToken(tokenString, referer)

		if !valid {
			ReturnUnauth(context)
		}

		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
			return
		}
		context.Keys["Username"] = claims.Username
	}
}
