package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/check", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
