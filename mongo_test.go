package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	Item   string `json:item bson:name`
	Qty    int32  `json:qty bson:size`
	Size   bson.D `json:size bson:size`
	Status string `json:status bson:status`
}

func TestItemQuery(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("order").Collection("item")

	items := []Item{}
	filter := bson.D{
		{"item", primitive.Regex{Pattern: "^p.*", Options: ""}},
	}
	fmt.Println(filter)
	filterOld := bson.D{
		{"item", bson.M{"$regex": "p"}},
	}
	fmt.Println(filterOld)
	cur, err := collection.Find(
		context.Background(),
		filterOld,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(items)
}
