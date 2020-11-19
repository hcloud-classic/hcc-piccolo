package graphqltype

import "github.com/graphql-go/graphql"

// VolumeAttachmentType : GraphQL type of VolumeAttachmentType
var VolumeAttachmentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VolumeAttachment",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"volume_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
