package main

import (
	"github.com/ProgramZheng/order-api/router"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Content-Type", "application/json")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 		c.Next()
// 	}
// }

func main() {
	// r := gin.Default()
	// r.Use(CORSMiddleware())
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run(":3000") // listen and serve on 0.0.0.0:8080
	r := router.Router

	router.SetRouter()

	r.Run(":3000")
}
