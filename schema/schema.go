package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ObjectID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ObjectID",
	Description: "The `primitive` scalar type represents a BSON Object.",
	// Serialize serializes `primitive.ObjectID` to string.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			v := *value
			return v.Hex()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `bson.ObjectId`.
	ParseValue: func(value interface{}) (result interface{}) {
		switch value := value.(type) {
		case string:
			result, _ = primitive.ObjectIDFromHex(value)
			return
		case *string:
			result, _ = primitive.ObjectIDFromHex(*value)
			return
		default:
			return nil
		}
		return nil
	},
	// ParseLiteral parses GraphQL AST to `bson.ObjectId`.
	ParseLiteral: func(valueAST ast.Value) (result interface{}) {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			result, _ = primitive.ObjectIDFromHex(valueAST.Value)
			return
		}
		return nil
	},
})

//定義Query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"hello":    &queryHello,
		"postById": &postById,
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
	Types:    []graphql.Type{ObjectID},
})
