package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Student struct {
	Name string
	Age  int
}

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
func insertOne(s Student) {
	initDB()
	collection := client.Database("go_db").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func insertMore(students []interface{}) {
	initDB()

	collection := client.Database("go_db").Collection("student")
	insertManyResult, err := collection.InsertMany(context.TODO(), students)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
func main() {
	s := Student{Name: "tom", Age: 20}
	//insertOne(s)

	s1 := Student{Name: "kite", Age: 21}
	s2 := Student{Name: "rose", Age: 22}
	students := []interface{}{s, s1, s2}
	insertMore(students)
}
