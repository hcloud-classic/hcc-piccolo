package graphqltype

import "github.com/graphql-go/graphql"

// QuotaType : Graphql object type of quota
var QuotaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Quota",
		Fields: graphql.Fields{
			"group_id": &graphql.Field{
				Type: graphql.Int,
			},
			"group_name": &graphql.Field{
				Type: graphql.String,
			},
			"limit_cpu_cores": &graphql.Field{
				Type: graphql.Int,
			},
			"limit_memory_gb": &graphql.Field{
				Type: graphql.Int,
			},
			"limit_subnet_host_bits": &graphql.Field{
				Type: graphql.Int,
			},
			"limit_adaptive_ip_cnt": &graphql.Field{
				Type: graphql.Int,
			},
			"pool_name": &graphql.Field{
				Type: graphql.String,
			},
			"limit_ssd_gb": &graphql.Field{
				Type: graphql.Int,
			},
			"limit_hdd_gb": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// QuotaListType : Graphql object type of QuotaList
var QuotaListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "QuotaList",
		Fields: graphql.Fields{
			"quota_list": &graphql.Field{
				Type: graphql.NewList(QuotaType),
			},
			"total_num": &graphql.Field{
				Type: graphql.Int,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// QuotaDetailType : Graphql object type of quotaDetail
var QuotaDetailType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "QuotaDetail",
		Fields: graphql.Fields{
			"group_id": &graphql.Field{
				Type: graphql.Int,
			},
			"group_name": &graphql.Field{
				Type: graphql.String,
			},
			"total_cpu_cores": &graphql.Field{
				Type: graphql.Int,
			},
			"total_memory_gb": &graphql.Field{
				Type: graphql.Int,
			},
			"nodes": &graphql.Field{
				Type: graphql.NewList(NodeType),
			},
			"total_subnet_host_bits": &graphql.Field{
				Type: graphql.Int,
			},
			"total_adaptive_ip_num": &graphql.Field{
				Type: graphql.Int,
			},
			"volumes": &graphql.Field{
				Type: graphql.NewList(VolumeType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
