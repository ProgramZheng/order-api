package main

import (
	"fmt"
	"testing"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Post struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
}

func TestPostQuery(t *testing.T) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	filter := bson.D{}
	post := Post{}
	_ = collection.FindOne(ctx, filter).Decode(&post)
	fmt.Println(post)
}
