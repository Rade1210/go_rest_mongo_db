package database

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseConnection struct {
	Client *mongo.Client
	Ctx context.Context
}

var db *DatabaseConnection

func New() *DatabaseConnection {
	if db == nil {
		client, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}
		return &DatabaseConnection{Client: client}
	}
	return db
}

func connect(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfuly")
	return nil
}