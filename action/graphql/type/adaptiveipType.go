package graphqltype

import "github.com/graphql-go/graphql"

// AdaptiveIPSettingType : Graphql object type of AdaptiveIPSetting
var AdaptiveIPSettingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIPSetting",
		Fields: graphql.Fields{
			"ext_ifaceip_address": &graphql.Field{
				Type: graphql.String,
			},
			"netmask": &graphql.Field{
				Type: graphql.String,
			},
			"gateway_address": &graphql.Field{
				Type: graphql.String,
			},
			"start_ip_address": &graphql.Field{
				Type: graphql.String,
			},
			"end_ip_address": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// AdaptiveIPAvailableIPListType : Graphql object type of AdaptiveIPAvailableIPList
var AdaptiveIPAvailableIPListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdaptiveIP",
		Fields: graphql.Fields{
			"available_ip_list": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
