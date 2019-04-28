package model

import (
	"fmt"
	"log"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Item struct {
	Item string
	qty  int32
	// size   bson.D
	status string
}

func (item *Item) Query(_query string) (items []Item, err error) {

	data := []Item{}
	client, ctx := mongodb.GetClient()
	collection := client.Database("order").Collection("item")

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

		item := Item{
			Item: result["item"].(string),
			qty:  result["qty"].(int32),
			// size:   result["size"].(bson.D),
			status: result["status"].(string),
		}
		data = append(data, item)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	items = data
	fmt.Println(data)

	return
}
