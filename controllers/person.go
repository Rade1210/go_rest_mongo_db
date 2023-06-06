package controllers

import (
	"net/http"
	"github.com/Rade1210/go_rest_mongo_db/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Id string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	Address Address `json:"address" bson:"address"`
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
}

type Address struct {
	AddressLine1 string `json:"addressLine1" bson:"addressLine1"`
	AddressLine2 string `json:"addressLine2" bson:"addressLine2"`
	City string `json:"city" bson:"city"`
	State string `json:"state" bson:"state"`
	ZipCode string `json:"zipCode" bson:"zipCode"`
	Country string `json:"country" bson:"country"`
}

func CreatePerson(ct *gin.Context){
	var person Person
		if err := ct.BindJSON(&person); err != nil {
			ct.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please provide the details in accepted format"})
		}
		SavePersonToDB(ct, person)
}

func GetPerson(ct *gin.Context) {
	id := ct.Param("id")
	db := database.New()
	collection := db.Client.Database("go_rest_mongo_db").Collection("Person")

	var result Person
	filter := bson.D{{"_id", id}}
	error := collection.FindOne(ct, filter).Decode(&result)
	if error != nil {
		ct.IndentedJSON(http.StatusNotFound, gin.H{"message":"Unable to find the person"})
	}

	ct.IndentedJSON(http.StatusOK, result)
}

func DeletePerson(ct *gin.Context) {
	id := ct.Param("id")
	db := database.New()
	collection := db.Client.Database("go_rest_mongo_db").Collection("Person")

	filter := bson.D{{"_id", id}}

	_, error := collection.DeleteOne(ct, filter)
	if error != nil {
		ct.IndentedJSON(http.StatusFailedDependency, gin.H{"message":"Unable to delete person from the records"})
	}

	ct.IndentedJSON(http.StatusAccepted, gin.H{"message":"Successfuly deleted the person"})
}

func UpdatePerson(ct *gin.Context) {
	id := ct.Param("id")
	db := database.New()
	collection := db.Client.Database("go_rest_mongo_db").Collection("Person")

	var person Person
	if err := ct.BindJSON(&person); err != nil {
		ct.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please provide the details in accepted format"})
	}

	filter := bson.M{"_id": id}
	updatePersonFirstName := bson.D{{
		Key: "$set", Value: bson.D{{Key: "firstName", Value: person.FirstName}},
	}}

	_, error := collection.UpdateOne(ct, filter, updatePersonFirstName)
	if error != nil {
		ct.IndentedJSON(http.StatusFailedDependency, gin.H{"message":"Unable to update the person in the records"})
	} else {
		ct.IndentedJSON(http.StatusAccepted, gin.H{"message":"Successfuly updated the person in the records"})
	}


}

func SavePersonToDB(ct *gin.Context, PersonRecord Person) {
	db := database.New()
	collection := db.Client.Database("go_rest_mongo_db").Collection("Person")
	response, insertError := collection.InsertOne(ct, PersonRecord)
	if insertError != nil {
		ct.IndentedJSON(http.StatusFailedDependency, gin.H{"message":"Unable to create Person in the database due to the internal failures"})
	}

	ct.IndentedJSON(http.StatusCreated, gin.H{"id": response.InsertedID})
} 