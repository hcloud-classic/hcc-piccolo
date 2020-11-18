<<<<<<< HEAD
package graphqlType
=======
package graphqltype
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33

import (
	"github.com/graphql-go/graphql"
)

// TelegrafType : GraphQL type of Telegraf
var TelegrafType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "telegraf",
		Fields: graphql.Fields{
<<<<<<< HEAD
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
				Type: graphql.NewList(SeriesType),
			},
		},
	},
)

// SeriesType : GraphQL type of Series
var SeriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "seriesType",
		Fields: graphql.Fields{
			"x": &graphql.Field{
				Type: graphql.Int,
			},
			"y": &graphql.Field{
				Type: graphql.Int,
=======
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
			},
		},
	},
)
