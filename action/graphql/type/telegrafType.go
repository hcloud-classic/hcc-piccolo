package graphqltype

import (
	"github.com/graphql-go/graphql"
)

// TelegrafType : GraphQL type of Telegraf
var TelegrafType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "telegraf",
		Fields: graphql.Fields{
			"metric": &graphql.Field{
				Type: graphql.String,
			},
			"subMetric": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: graphql.NewList(graphql.NewList(graphql.Float)),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// SeriesType : GraphQL type of Series
var SeriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "seriesType",
		Fields: graphql.Fields{
			"values": &graphql.Field{
				Type: graphql.NewList(graphql.Float),
			},
		},
	},
)
