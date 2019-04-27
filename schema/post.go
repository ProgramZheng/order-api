package schema

import (
	"fmt"

	"github.com/ProgramZheng/order-api/model"

	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Post",
	Description: "Post Model",
	Fields: graphql.Fields{
		"_query": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryPost = graphql.Field{
	Name:        "QueryPost",
	Description: "Query Post",
	Type:        graphql.NewList(postType),

	Args: graphql.FieldConfigArgument{
		"_query": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		_query, _ := p.Args["_query"].(string)

		id, _ := p.Args["id"].(int)
		title, _ := p.Args["title"].(string)

		fmt.Println("/" + title + "$/")
		return (&model.Post{}).Query(_query, id, title)
	},
}
