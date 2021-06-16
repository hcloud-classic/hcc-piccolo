package dao

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
	"strings"
)

// ReadQuota : Get the quota of the group
func ReadQuota(groupID int64) (*pb.GroupQuota, error) {
	var quota pb.GroupQuota

	var groupName string
	var limitCPUCores int
	var limitMemoryGB int
	var limitSubnetCnt int
	var limitAdaptiveIPCnt int
	var poolName string
	var limitSSDGB int
	var limitHDDGB int

	sql := "select piccolo.quota.group_id, piccolo.group.name as group_name, " +
		"piccolo.quota.limit_cpu_cores, piccolo.quota.limit_memory_gb, " +
		"piccolo.quota.limit_subnet_cnt, piccolo.quota.limit_adaptive_ip_cnt, " +
		"piccolo.quota.pool_name, piccolo.quota.limit_ssd_gb, piccolo.quota.limit_hdd_gb" +
		" from piccolo.quota, piccolo.group where piccolo.quota.group_id = piccolo.group.id" +
		" and piccolo.quota.group_id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row,
		&groupID,
		&groupName,
		&limitCPUCores,
		&limitMemoryGB,
		&limitSubnetCnt,
		&limitAdaptiveIPCnt,
		&poolName,
		&limitSSDGB,
		&limitHDDGB)
	if err != nil {
		errStr := "ReadGroup(): " + err.Error()
		logger.Logger.Println(errStr)

		return nil, err
	}

	quota = pb.GroupQuota{
		GroupID:            groupID,
		GroupName:          groupName,
		LimitCPUCores:      int32(limitCPUCores),
		LimitMemoryGB:      int32(limitMemoryGB),
		LimitSubnetCnt:     int32(limitSubnetCnt),
		LimitAdaptiveIPCnt: int32(limitAdaptiveIPCnt),
		PoolName:           poolName,
		LimitSSDGB:         int32(limitSSDGB),
		LimitHDDGB:         int32(limitHDDGB),
	}

	return &quota, nil
}

// CreateQuota : Create the quota of the group
func CreateQuota(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(int)
	poolName, poolNameOk := args["pool_name"].(string)
	ssdSize, ssdSizeOk := args["ssd_size"].(int)
	hddSize, hddSizeOk := args["hdd_size"].(int)
	selectedNodes, selectedNodesOk := args["selected_nodes"].(string)
	subnetCnt, subnetCntOk := args["subnet_cnt"].(int)
	adaptiveCnt, adaptiveCntOk := args["adaptive_cnt"].(int)

	if isMaster && !groupIDOk {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need a group_id argument")}, nil
	}

	if !poolNameOk || !ssdSizeOk || !hddSizeOk || !selectedNodesOk ||
		!subnetCntOk || !adaptiveCntOk {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need pool_name and ssd_size, hdd_size, selected_nodes, subnet_cnt, adaptive_cnt arguments")}, nil
	}

	selectedNodesSplit := strings.Split(selectedNodes, ",")
	if len(selectedNodesSplit) == 0 {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"nodes are not selected or invalid string is provided for selected_nodes argument")}, nil
	}

	var limitCPUCores = 0
	var limitMemoryGB = 0

	for _, nodeUUID := range selectedNodesSplit {
		resGetNode, err := client.RC.GetNode(nodeUUID)
		if err != nil {
			return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError,
				"Failed to get the node's info (nodeUUID="+nodeUUID+")")}, nil
		}

		limitCPUCores += int(resGetNode.Node.CPUCores)
		limitMemoryGB += int(resGetNode.Node.Memory)
	}

	resPoolList, err := client.RC.GetPoolList(&pb.ReqGetPoolList{
		Pool: &pb.Pool{
			Action: "read",
		},
	})
	if err != nil {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError,
			"Failed to get pool list")}, nil
	}

	var poolNameFound = false
	for _, pool := range resPoolList.Pool {
		if pool.Name == poolName {
			poolNameFound = true
			break
		}
	}

	if !poolNameFound {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"Provided pool name is not exist in the pool list")}, nil
	}

	group, err := ReadGroup(groupID)
	if err != nil {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError,
			"Failed to get the group's info")}, nil
	}

	for _, nodeUUID := range selectedNodesSplit {
		_, err := client.RC.UpdateNode(&pb.ReqUpdateNode{
			Node: &pb.Node{
				UUID:    nodeUUID,
				GroupID: int64(groupID),
			},
		})
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError,
				"Failed to insert groupID to the node (nodeUUID="+nodeUUID+")")}, nil
		}
	}

	quota := model.Quota{
		GroupID:            int64(groupID),
		GroupName:          group.Name,
		LimitCPUCores:      limitCPUCores,
		LimitMemoryGB:      limitMemoryGB,
		LimitSubnetCnt:     subnetCnt,
		LimitAdaptiveIPCnt: adaptiveCnt,
		PoolName:           poolName,
		LimitSSDGB:         ssdSize,
		LimitHDDGB:         hddSize,
	}

	sql := "insert into quota(group_id, limit_cpu_cores, limit_memory_gb, limit_subnet_cnt, limit_adaptive_ip_cnt, pool_name, limit_ssd_gb, limit_hdd_gb) values (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "CreateQuota(): " + err.Error()
		logger.Logger.Println(errStr)

		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(quota.GroupID, quota.LimitCPUCores, quota.LimitMemoryGB, quota.LimitSubnetCnt, quota.LimitAdaptiveIPCnt, quota.PoolName, quota.LimitSSDGB, quota.LimitHDDGB)
	if err != nil {
		errStr := "CreateQuota(): " + err.Error()
		logger.Logger.Println(errStr)

		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	return &quota, nil
}
