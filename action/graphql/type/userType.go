package graphqltype

import "github.com/graphql-go/graphql"

// UserType : Graphql object type of user
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"login_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// Token : Graphql object type of token
var Token = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// IsValid : Graphql object type of IsValid
var IsValid = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IsValid",
		Fields: graphql.Fields{
			"isvalid": &graphql.Field{
				Type: graphql.Boolean,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
