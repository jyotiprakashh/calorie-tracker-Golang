package routes

import (
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func DBinstances() *mongo.Client {
	MongoDb := "mongodb://localhost:27017/caloriesdb"
	// client, err := mongo.NewClient(options.Clients().ApplyURI(MongoDb))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		panic(err)
		fmt.Print("error occured while connecting to mongodb")
	}

	cts, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = client.Ping(cts, nil)
		if err != nil {
		panic(err)
		fmt.Print("error occured while connecting to mongodb")
	}
	fmt.Print("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstances()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName)
	return collection
}