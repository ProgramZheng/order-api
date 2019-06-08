package schema

import (
	"fmt"

	"github.com/ProgramZheng/order-api/model"
	"go.mongodb.org/mongo-driver/bson"

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
			"_id": &graphql.Field{
				Type: ObjectID,
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

var postById = graphql.Field{
	Name:        "Post By Id",
	Description: "依照id取得Post",
	Type:        postType,

	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: ObjectID,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		filter := bson.D{bson.E{"_id", params.Args["_id"]}}
		model, err := model.ById(filter)
		fmt.Println(model)
		//
		return model, err
	},
}

var postList = graphql.Field{
	Name:        "postList",
	Description: "依照id或title取得Post",
	Type:        graphql.NewList(postType),

	Args: graphql.FieldConfigArgument{
		// "_query": &graphql.ArgumentConfig{
		// 	Type: graphql.String,
		// },
		"_id": &graphql.ArgumentConfig{
			Type: ObjectID,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	//接到請求後，執行的函數
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		type result struct {
			data interface{}
			err  error
		}
		//
		filter := bson.D{}
		for key, value := range params.Args {
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
	Description: "新增Post",
	Type:        postType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"text": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		data := bson.M{
			"title": params.Args["title"],
			"text":  params.Args["text"],
		}
		model, err := model.Add(data)
		//
		return model, err
	},
}
