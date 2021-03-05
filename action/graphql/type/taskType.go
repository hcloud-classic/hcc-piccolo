package graphqltype

import "github.com/graphql-go/graphql"

// TaskNoRecursiveType : Graphql object type of TaskNoRecursive
var TaskNoRecursiveType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TaskNoRecursiveType",
		Fields: graphql.Fields{
			"cmd": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"pid": &graphql.Field{
				Type: graphql.Int,
			},
			"ppid": &graphql.Field{
				Type: graphql.Int,
			},
			"pgid": &graphql.Field{
				Type: graphql.Int,
			},
			"sid": &graphql.Field{
				Type: graphql.Int,
			},
			"priority": &graphql.Field{
				Type: graphql.Int,
			},
			"nice": &graphql.Field{
				Type: graphql.Int,
			},
			"num_threads": &graphql.Field{
				Type: graphql.Int,
			},
			"start_time": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_usage": &graphql.Field{
				Type: graphql.String,
			},
			"mem_usage": &graphql.Field{
				Type: graphql.String,
			},
			"epm_type": &graphql.Field{
				Type: graphql.String,
			},
			"epm_source": &graphql.Field{
				Type: graphql.Int,
			},
			"epm_target": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// TaskType : Graphql object type of Task
var TaskType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TaskType",
		Fields: graphql.Fields{
			"cmd": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"pid": &graphql.Field{
				Type: graphql.Int,
			},
			"ppid": &graphql.Field{
				Type: graphql.Int,
			},
			"pgid": &graphql.Field{
				Type: graphql.Int,
			},
			"sid": &graphql.Field{
				Type: graphql.Int,
			},
			"priority": &graphql.Field{
				Type: graphql.Int,
			},
			"nice": &graphql.Field{
				Type: graphql.Int,
			},
			"num_threads": &graphql.Field{
				Type: graphql.Int,
			},
			"children": &graphql.Field{
				Type: graphql.NewList(TaskNoRecursiveType),
			},
			"threads": &graphql.Field{
				Type: graphql.NewList(TaskNoRecursiveType),
			},
			"start_time": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_usage": &graphql.Field{
				Type: graphql.String,
			},
			"mem_usage": &graphql.Field{
				Type: graphql.String,
			},
			"epm_type": &graphql.Field{
				Type: graphql.String,
			},
			"epm_source": &graphql.Field{
				Type: graphql.Int,
			},
			"epm_target": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// TaskListType : Graphql object type of TaskList
var TaskListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TaskList",
		Fields: graphql.Fields{
			"task_list": &graphql.Field{
				Type: graphql.NewList(TaskType),
			},
			"total_tasks": &graphql.Field{
				Type: graphql.Int,
			},
			"total_mem_usage": &graphql.Field{
				Type: graphql.String,
			},
			"total_mem": &graphql.Field{
				Type: graphql.String,
			},
			"total_mem_usage_percent": &graphql.Field{
				Type: graphql.String,
			},
			"total_cpu_usage": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)
