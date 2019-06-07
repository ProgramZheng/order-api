package schema

import (
	"github.com/graphql-go/graphql"
)

//定義跟查詢節點
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello":    &queryHello,
		"postList": &postList,
		"item":     &queryItem,
	},
})

//定義Schema用於http handler處理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, //需透過GraphQL更新數據所使用的
})
