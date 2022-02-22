package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// TimapniService : GraphQL type of TimapniService
var TimapniService = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TimapniService",
		Fields: graphql.Fields{
			"target": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: ErrorField,
			},
		},
	},
)

// TimpaniMasterSync : GraphQL type of TimpaniMasterSync
var TimpaniMasterSync = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TimpaniMasterSync",
		Fields: graphql.Fields{
			"runstatus": &graphql.Field{
				Type: graphql.String,
			},
			"runuuid": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: ErrorField,
			},
		},
	},
)

// RestoreType : Graphql object type of restore
var RestoreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Restore",
		Fields: graphql.Fields{
			"runstatus": &graphql.Field{
				Type: graphql.String,
			},
			"runuuid": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: ErrorField,
			},
		},
	},
)

// ErrorField : Graphql object type of errors
var ErrorField = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ErrorField",
		Fields: graphql.Fields{
			"errcode": &graphql.Field{
				Type: graphql.String,
			},
			"errmsg": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
