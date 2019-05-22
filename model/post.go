package model

import (
	"fmt"
	"log"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Post struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
}

func Query(filter bson.D) (posts []interface{}, err error) {

	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	// fmt.Println(filter)
	cur, err := collection.Find(ctx, filter)
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
		fmt.Println(post)
		//return data
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
