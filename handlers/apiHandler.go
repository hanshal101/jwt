package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAPIHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"access": "true"})
}
