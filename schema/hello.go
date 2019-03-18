package schema

import (
	"github.com/ProgramZheng/order-api/model"

	"github.com/graphql-go/graphql"
)

//定義查詢的對象，支持嵌套
var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

//處理查詢請求
var queryHello = graphql.Field{
	Name:        "QueryHello",
	Description: "Query Hello",
	Type:        graphql.NewList(helloType),

	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {

		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)

		return (&model.Hello{}).Query(id, name)
	},
}
