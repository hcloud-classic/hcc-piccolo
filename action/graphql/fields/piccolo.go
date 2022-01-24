package fields

import (
	"hcc/piccolo/action/graphql/mutationparser"
	"hcc/piccolo/action/graphql/queryparser"
	"hcc/piccolo/action/graphql/queryparserext"
	graphqlType "hcc/piccolo/action/graphql/type"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"

	"github.com/graphql-go/graphql"
	"innogrid.com/hcloud-classic/hcc_errors"
)

var Login = graphql.Field{
	Type:        graphqlType.Token,
	Description: "Execute login process for piccolo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		data, err := queryparser.Login(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / login: " + err.Error())
		}
		return data, err
	},
}

var User = graphql.Field{
	Type:        graphqlType.UserType,
	Description: "Get the user info from piccolo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparserext.User(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / user: " + err.Error())
		}
		return data, err
	},
}

var ListUser = graphql.Field{
	Type:        graphqlType.UserListType,
	Description: "Get the user list from piccolo",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"authentication": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.UserList(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / list_user: " + err.Error())
		}
		return data, err
	},
}

var NumUser = graphql.Field{
	Type:        graphqlType.UserNumType,
	Description: "Get the number of users",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NumUser(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / num_user: " + err.Error())
		}
		return data, err
	},
}

var AllGroup = graphql.Field{
	Type:        graphqlType.GroupListType,
	Description: "Get the group list from piccolo",
	Args: graphql.FieldConfigArgument{
		"include_master": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.GroupList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ReadGroupList(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("piccolo / all_group: " + err.Error())
		}
		return data, err
	},
}

var CheckToken = graphql.Field{
	Type:        graphqlType.IsValid,
	Description: "Check validation of the token for piccolo",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		data, err := queryparser.CheckToken(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / check_token: " + err.Error())
		}
		return data, err
	},
}

var NodeAvailable = graphql.Field{
	Type:        graphqlType.NodeAvailableType,
	Description: "Get available info of nodes",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.NodeAvailable{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.NodeAvailable(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / node_available: " + err.Error())
		}
		return data, err
	},
}

var ServerLog = graphql.Field{
	Type:        graphqlType.ServerActionsType,
	Description: "Get the server's log",
	Args: graphql.FieldConfigArgument{
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ServerActions{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := dao.ShowServerActions(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / server_log: " + err.Error())
		}
		return data, err
	},
}

var ServerAlarmList = graphql.Field{
	Type:        graphqlType.ServerAlarmsType,
	Description: "Get the server's alarm list",
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_uuid": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"server_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"reason": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ServerAlarms{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := dao.ShowServerAlarms(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / server_alarm_list: " + err.Error())
		}
		return data, err
	},
}

var NumUnreadServerAlarm = graphql.Field{
	Type:        graphqlType.ServerAlarmsNumType,
	Description: "Get the number of unread server's alarm",
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, _, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ServerAlarmsNum{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := dao.ShowUnreadServerAlarmsNum(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / num_unread_server_alarm: " + err.Error())
		}
		return data, err
	},
}

var Quota = graphql.Field{
	Type:        graphqlType.QuotaListType,
	Description: "Get info of the quota from piccolo",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.ReadQuota(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / quota: " + err.Error())
		}
		return data, err
	},
}

var ListQuota = graphql.Field{
	Type:        graphqlType.QuotaListType,
	Description: "Get the quota list from piccolo",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"total_cpu_cores": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"total_memory_gb": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit_subnet_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit_adaptive_ip_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit_node_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"pool_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"limit_ssd_gb": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit_hdd_gb": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.QuotaList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.QuotaList(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / list_quota: " + err.Error())
		}
		return data, err
	},
}

var QuotaDetail = graphql.Field{
	Type:        graphqlType.QuotaDetailType,
	Description: "Get detail info of the quota from piccolo",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := queryparser.QuotaDetail(params.Args, isAdmin, isMaster)
		if err != nil {
			logger.Logger.Println("piccolo / quota_detail: " + err.Error())
		}
		return data, err
	},
}

var SignUp = graphql.Field{
	Type:        graphqlType.UserType,
	Description: "Execute user sign up process for piccolo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"authentication": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := mutationparser.SignUp(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / signup: " + err.Error())
		}
		return data, err
	},
}

var Unregister = graphql.Field{
	Type:        graphqlType.UserType,
	Description: "Execute user unregister process for piccolo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.Unregister(params.Args, isAdmin, isMaster, id, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / unregister: " + err.Error())
		}
		return data, err
	},
}

var UpdateUser = graphql.Field{
	Type:        graphqlType.UserType,
	Description: "Update user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"authentication": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.UpdateUser(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / update_user: " + err.Error())
		}
		return data, err
	},
}

var CreateUser = graphql.Field{
	Type:        graphqlType.GroupType,
	Description: "Get the group info",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.CreateGroup(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("piccolo / create_group: " + err.Error())
		}
		return data, err
	},
}

var UpdateGroup = graphql.Field{
	Type:        graphqlType.GroupType,
	Description: "Update the group info",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.UpdateGroup(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / update_group: " + err.Error())
		}
		return data, err
	},
}

var DeleteGroup = graphql.Field{
	Type:        graphqlType.GroupType,
	Description: "Delete the group info",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, _, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := mutationparser.DeleteGroup(params.Args, isMaster)
		if err != nil {
			logger.Logger.Println("piccolo / delete_group: " + err.Error())
		}
		return data, err
	},
}

var CreateGroup = graphql.Field{
	Type:        graphqlType.QuotaType,
	Description: "Create quota",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"pool_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"ssd_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"hdd_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"subnet_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"adaptive_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"node_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := dao.CreateQuota(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / create_quota: " + err.Error())
		}
		return data, err
	},
}

var UpdateQuota = graphql.Field{
	Type:        graphqlType.QuotaType,
	Description: "Update quota",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"pool_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"ssd_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"hdd_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"subnet_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"adaptive_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"node_cnt": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := dao.UpdateQuota(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / update_quota: " + err.Error())
		}
		return data, err
	},
}

var DeleteQuota = graphql.Field{
	Type:        graphqlType.QuotaType,
	Description: "Delete quota",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		isAdmin, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, true)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := dao.DeleteQuota(params.Args, isAdmin, isMaster, int(groupID))
		if err != nil {
			logger.Logger.Println("piccolo / delete_quota: " + err.Error())
		}
		return data, err
	},
}

var DeleteServerAlarm = graphql.Field{
	Type:        graphqlType.ServerAlarmType,
	Description: "Delete the server alarm",
	Args: graphql.FieldConfigArgument{
		"no": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, _, id, _, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		data, err := dao.DeleteServerAlarm(params.Args, id)
		if err != nil {
			logger.Logger.Println("piccolo / delete_group: " + err.Error())
		}
		return data, err
	},
}

// SubScription

var ListUser = graphql.Field{
	Type:        graphqlType.UserListType,
	Description: "Get the user list from piccolo",
	Args: graphql.FieldConfigArgument{
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"authentication": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"row": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.UserList(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / list_user (Subscription): " + err.Error())
		}
		return data, err
	},
}

var ResourceUsage = graphql.Field{
	Type:        graphqlType.ResourceUsageType,
	Description: "Get resource usage",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		_, isMaster, _, groupID, err := usertool.ValidateToken(params.Args, false)
		if err != nil {
			return model.ResourceUsage{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, err.Error())}, nil
		}
		if !isMaster {
			params.Args["group_id"] = int(groupID)
		}
		data, err := queryparser.ResourceUsage(params.Args)
		if err != nil {
			logger.Logger.Println("piccolo / resource_usage (Subscription): " + err.Error())
		}
		return data, err
	},
}
