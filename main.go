package main

import (
	"github.com/Rade1210/go_rest_mongo_db/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Address: localhost:8080
	router.POST("/postPerson", controllers.CreatePerson)
	router.GET("/getPerson/:id", controllers.GetPerson)
	router.DELETE("/deletePerson/:id", controllers.DeletePerson)
	router.Run()
}