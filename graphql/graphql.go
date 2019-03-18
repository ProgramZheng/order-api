package graphql

import (
	"github.com/ProgramZheng/order-api/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	//透過gin進行封裝
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
