package main

import (
	"context"
	"fmt"
	"log"

	"github.com/anhtuanqn1002/redis" // packet gif dau :)) go module link dau

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	redisClient := redis.RClient()

	var name = "name"
	value, err := redis.Get(redisClient, name)
	if err == nil {
		fmt.Println(value)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mydb").Collection("persons")
	insertResult, err := collection.InsertOne(context.TODO(), bson.D{{"name", value}})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	filter := bson.D{}
	result := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.TODO()) {
		elem := &bson.D{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		// ideally, you would do something with elem....
		// but for now just print it to the console
		fmt.Println(elem)
	}
	fmt.Println(result)
}
