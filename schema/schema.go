package schema

import (
	"github.com/graphql-go/graphql"
)

//定義Query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello":    &queryHello,
		"postById": postById,
		"postList": &postList,
		"item":     &queryItem,
	},
})

//定義Mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootMutation",
	Description: "Root Mutation",
	Fields: graphql.Fields{
		"addPost": &addPost,
	},
})

//定義Schema用於http handler處理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation, //需透過GraphQL更新數據所使用的
})
