package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func initDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
}

func del() {
	c := client.Database("go_db").Collection("student")
	ctx := context.TODO()

	dr, err := c.DeleteMany(ctx, bson.D{{"name", "tom"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}
func main() {
	initDB()
	del()
}
