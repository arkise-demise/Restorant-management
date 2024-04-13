package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-mongodb.org/mongo-driver/mongo"
	"go-mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	MongoDb := "mongodb://localhost:27027"
	fmt.Print(MongoDb)

	client,err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)

	}
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)

	fmt.Println("connected to mongodb")

	return client

	}
}

var Client *mongo.Client =DBinstance()

func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	var colllection *mongo.collection = client.Database("restaurant").collection(collectionName)

	return collection
}