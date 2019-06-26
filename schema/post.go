package schema

import (
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

var postNestedInputObject = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"text": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// var someInputObject = graphql.NewInputObject(
// 	graphql.InputObjectConfig{
// 		Name: "SomeInputObject",
// 		Fields: graphql.InputObjectConfigFieldMap{
// 			"nested": &graphql.InputObjectFieldConfig{
// 				Type: postNestedInputObject,
// 			},
// 		},
// 	},
// )

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
		model, err := model.ByID(filter)
		//
		return model, err
	},
}

var postList = graphql.Field{
	Name: "postList",
	Description: `
	query{
		postList(
			title: title,
			text: text
		) {
			_id
			text
			title
		}
	}
	`,
	Type: graphql.NewList(postType),

	Args: graphql.FieldConfigArgument{
		// "_query": &graphql.ArgumentConfig{
		// 	Type: graphql.String,
		// },
		// "_id": &graphql.ArgumentConfig{
		// 	Type: ObjectID,
		// },
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
	Name: "addPost",
	Description: `
	mutation{
		addPost(
			title:title,
			text:text
		) {
			_id
			text
			title
		}
	}
	`,
	Type: postType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"text": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: "",
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

var addManyPost = graphql.Field{
	Name: "addManyPost",
	Description: `
	mutation{
		addManyPost(
			array:[
				{title:title},
				{title:title,text:text},
			],
		) {
			_id
			text
			title
		}
	}
	`,
	Type: postType,
	Args: graphql.FieldConfigArgument{
		"array": &graphql.ArgumentConfig{
			Type: graphql.NewList(postNestedInputObject),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		data := []params.Args["array"]
		fmt.Println(data)
		// for key, value := range array {
		// 	data = append(data, bson.M{key: value})
		// }

		// model, err := model.AddMany(params.Args["array"])
		//
		// return interface{}, error
		return model, err
	},
}

var updateOnePost = graphql.Field{
	Name: "updateOnePost",
	Description: `
	mutation{
		updateOnePost(
			_id:_id,
			title:title,
		) {
			_id
			text
			title
		}
	}`,
	Type: postType,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(ObjectID),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"text": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: "",
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		filter := bson.D{bson.E{"_id", params.Args["_id"]}}
		update := bson.M{
			"$set": bson.M{
				"title": params.Args["title"],
				"text":  params.Args["text"],
			},
		}
		model, err := model.UpdateOne(filter, update)
		if err != nil {
			// log.Fatal(err)
		}
		//
		return model, err
	},
}

var deleteOnePost = graphql.Field{
	Name: "deleteOnePost",
	Description: `
	mutation{
		deleteOnePost(
			_id:_id
		) {
			_id
			text
			title
		}
	}`,
	Type: postType,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(ObjectID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		filter := bson.D{bson.E{"_id", params.Args["_id"]}}
		model, err := model.DeleteOne(filter)
		if err != nil {
			// log.Fatal(err)
		}
		//
		return model, err
	},
}
