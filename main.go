package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ProgramZheng/order-api/mongodb"
	"github.com/ProgramZheng/order-api/router"
	"github.com/mongodb/mongo-go-driver/bson"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Content-Type", "application/json")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 		c.Next()
// 	}
// }

func createData() {
	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}
	client, _ := mongodb.GetClient()
	collection := client.Database("order").Collection("item")
	result, _ := collection.InsertMany(context.Background(), docs)
	fmt.Println(result)
}

func main() {
	// r := gin.Default()
	// r.Use(CORSMiddleware())
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run(":3000") // listen and serve on 0.0.0.0:8080

	// createData()
	fmt.Println(os.Getenv("DB_HOST"))
	r := router.Router

	router.SetRouter()

	r.Run(":3000")
}
