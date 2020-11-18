package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// VncPortType : Graphql object type of ControlVncType
var VncPortType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VncPort",
		Fields: graphql.Fields{
			"port": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
