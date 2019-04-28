package schema

import (
	"github.com/ProgramZheng/order-api/model"

	"github.com/graphql-go/graphql"
)

var itemType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Item",
	Description: "Item Model",
	Fields: graphql.Fields{
		"_query": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: graphql.String,
		},
		"qty": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryItem = graphql.Field{
	Name:        "QueryItem",
	Description: "Query Item",
	Type:        graphql.NewList(itemType),

	Args: graphql.FieldConfigArgument{
		"_query": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"item": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"qty": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		_query, _ := p.Args["_query"].(string)

		return (&model.Item{}).Query(_query)
	},
}
