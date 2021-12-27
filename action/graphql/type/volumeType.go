package graphqltype

import "github.com/graphql-go/graphql"

// VolumeType : GraphQL type of VolumeType
var VolumeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Volume",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"size": &graphql.Field{
				Type: graphql.Int,
			},
			"group_id": &graphql.Field{
				Type: graphql.Int,
			},
			"filesystem": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"network_ip": &graphql.Field{
				Type: graphql.String,
			},
			"use_type": &graphql.Field{
				Type: graphql.String,
			},
			"user_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"lun_num": &graphql.Field{
				Type: graphql.Int,
			},
			"pool": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// VolumeListType : Graphql object type of ServerList
var VolumeListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VolumeList",
		Fields: graphql.Fields{
			"volume_list": &graphql.Field{
				Type: graphql.NewList(VolumeType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
