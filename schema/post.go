package schema

import (
	"github.com/ProgramZheng/order-api/model"

	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Post",
	Description: "Post Model",
	Fields: graphql.Fields{
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
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

		id, _ := p.Args["id"].(int)
		title, _ := p.Args["title"].(string)

		return (&model.Post{}).Query(id, title)
	},
}
