package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Address: localhost:8080
	
	router.GET("/ping", func(ct *gin.Context) {
		ct.IndentedJSON(http.StatusOK, gin.H{"message": "Ping successful"})
	})

	router.Run()
}