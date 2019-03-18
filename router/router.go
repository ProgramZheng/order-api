package router

import (
	"github.com/ProgramZheng/order-api/graphql"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	//GET方式提供GraphQL Web介面的操作
	//可依照自己需求更改POST或GET
	Router.POST("/graphql", graphql.GraphqlHandler())
	Router.GET("/graphql", graphql.GraphqlHandler())
}
