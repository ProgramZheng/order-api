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

func (post *Post) Query(id int, title string) (posts []Post, err error) {

	data := []Post{}
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	if id == 0 && title == "" {
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
			data = append(data, post)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		posts = data
	}

	result := Post{}
	filter := bson.M{}
	if id > 0 {
		filter = bson.M{"id": id}
	}

	if title != "" {
		filter = bson.M{"title": title}
	}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// for _, v := range data {
	// 	// posts = append(posts, v)
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
