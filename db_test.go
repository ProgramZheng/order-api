package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Post struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
}

func TestQuery(t *testing.T) {
	posts := []Post{}
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)

		post := Post{
			Id:    result["id"].(int32),
			Title: result["title"].(string),
		}
		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
}
