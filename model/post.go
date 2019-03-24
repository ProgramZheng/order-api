package model

import (
	"log"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Post struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
}

func (post *Post) Query(id int, title string) (posts []Post, err error) {
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

		post := Post{
			Id:    result["id"].(int32),
			Title: result["title"].(string),
		}
		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// for _, v := range result {
	// 	if id >= 0 && len(title) <= 0 {
	// 		if v.Id == id {
	// 			posts = append(posts, v)
	// 		}
	// 	}

	// 	if id < 0 && len(title) > 0 {
	// 		if v.Title == title {
	// 			posts = append(posts, v)
	// 		}
	// 	}

	// 	if id >= 0 && len(title) > 0 {
	// 		if v.Title == title && v.Id == id {
	// 			posts = append(posts, v)
	// 		}
	// 	}

	// 	if id < 0 && len(title) <= 0 {
	// 		posts = append(posts, v)
	// 	}
	// }
	return
}