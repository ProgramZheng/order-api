package model

import (
	"fmt"
	"log"

	"github.com/ProgramZheng/order-api/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
	Text  string             `json:"text" bson:"text"`
}

func ById(filter bson.D) (post Post, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func List(filter bson.D) (posts []interface{}, err error) {

	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	cur, err := collection.Find(ctx, filter)
	fmt.Println(filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var post = Post{}
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		//return data
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func Add(data bson.M) (id interface{}, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	res, err := collection.InsertOne(ctx, data)
	id = res.InsertedID
	if err != nil {
		log.Fatal(err)
	}
	return
}
