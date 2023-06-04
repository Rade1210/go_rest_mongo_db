package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Address Address `json:"address"`
	EmailAddress string `json:"emailAddress"`
}

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}

func main() {
	router := gin.Default() // Address: localhost:8080
	router.POST("/postPerson", func(ct *gin.Context) {
		var person Person
		if err := ct.BindJSON(&person); err != nil {
			ct.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please provide the details in accepted format"})
		}
		fmt.Println(person)
	})
	router.Run()
}