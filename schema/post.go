package schema

import (
	"fmt"

	"github.com/ProgramZheng/order-api/model"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Post",
		Description: "Post Model",
		Fields: graphql.Fields{
			// "_query": &graphql.Field{
			// 	Type: graphql.String,
			// },
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var postById = &graphql.Field{
	Name:        "Post By Id",
	Description: "Get post to once",
	Type:        postType,

	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//
		filter := bson.D{bson.E{"id", p.Args["id"]}}
		model, err := model.ById(filter)
		fmt.Println(model)
		//
		return model, err
	},
}

var postList = graphql.Field{
	Name:        "postList",
	Description: "Get post to list",
	Type:        graphql.NewList(postType),

	Args: graphql.FieldConfigArgument{
		// "_query": &graphql.ArgumentConfig{
		// 	Type: graphql.String,
		// },
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		type result struct {
			data interface{}
			err  error
		}
		//
		filter := bson.D{}
		for key, value := range p.Args {
			switch value.(type) {
			case int:
				filter = append(filter, bson.E{key, value})
			case string:
				filter = append(filter, bson.E{key, bson.M{"$regex": value}})
			}
		}
		model, err := model.List(filter)
		//
		ch := make(chan *result, 1)
		go func() {
			ch <- &result{data: model, err: err}
		}()
		return func() (interface{}, error) {
			r := <-ch
			return r.data, r.err
		}, nil
	},
}

var addPost = graphql.Field{
	Name:        "postList",
	Description: "Get post to list",
	Type:        graphql.NewList(postType),
}
