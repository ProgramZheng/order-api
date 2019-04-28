package schema

import (
	"github.com/ProgramZheng/order-api/model"
	"github.com/mongodb/mongo-go-driver/bson"

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
		filter := bson.D{}
		for key, value := range p.Args {
			switch value.(type) {
			case int:
				filter = append(filter, bson.E{key, value})
			case string:
				filter = append(filter, bson.E{key, bson.M{"$regex": value}})
			}
		}
		// fmt.Println(filter)
		return (&model.Post{}).Query(filter)
	},
}
