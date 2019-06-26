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

//ByID get post for id
func ByID(filter bson.D) (post Post, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	err = collection.FindOne(ctx, filter).Decode(&post)
	return
}

//List get post list
func List(filter bson.D) (posts []interface{}, err error) {

	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	cur, err := collection.Find(ctx, filter)

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var post = Post{}
		fmt.Println(cur)
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

// Add post
func Add(data bson.M) (post Post, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	res, err := collection.InsertOne(ctx, data)
	id := res.InsertedID.(primitive.ObjectID)
	//Add return result ID to post struct
	post.ID = id
	//data(bson.M) to bytes
	bsonBytes, _ := bson.Marshal(data)
	//bytes to post struct
	bson.Unmarshal(bsonBytes, &post)
	return
}

// AddMany post
func AddMany(data []interface{}) (res interface{}, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	res, err = collection.InsertMany(ctx, data)
	// var IDs interface{}
	// IDs = res.InsertedIDs
	fmt.Println(res)
	//Add return result ID to post struct
	// post.ID = id
	// //data(bson.M) to bytes
	// bsonBytes, _ := bson.Marshal(data)
	// //bytes to post struct
	// bson.Unmarshal(bsonBytes, &post)
	return
}

// UpdateOne post
func UpdateOne(filter bson.D, update bson.M) (post Post, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	//Get post
	collection.FindOneAndUpdate(ctx, filter, update).Decode(&post)
	return
}

// DeleteOne post
func DeleteOne(filter bson.D) (post Post, err error) {
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("post")
	//Get post
	collection.FindOneAndDelete(ctx, filter).Decode(&post)

	return
}
