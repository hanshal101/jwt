package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(res *gin.Context) {
	res.SetCookie("apikey", "", 0, "/", "localhost:9876", true, true)
	res.JSON(http.StatusAccepted, gin.H{"message": "Logged Out"})
}
