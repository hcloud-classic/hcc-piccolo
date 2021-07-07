package dao

import (
	"errors"
	"hcc/piccolo/action/graphql/queryparserext"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
	"strconv"
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
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("quota is not exist")
		}

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

func checkPoolSize(quota model.Quota, resGetServerList *pb.ResGetServerList) error {
	if resGetServerList == nil {
		_resGetServerList, err := client.RC.GetServerList(
			&pb.ReqGetServerList{
				Server: &pb.Server{
					GroupID: quota.GroupID,
				},
			})
		if err != nil {
			return errors.New("failed to get server list")
		}

		resGetServerList = _resGetServerList
	}

	resPoolList, err := client.RC.GetPoolList(&pb.ReqGetPoolList{
		Pool: &pb.Pool{
			Action: "read",
		},
	})
	if err != nil {
		return errors.New("failed to get pool list")
	}

	var volumesSize int
	for _, server := range resGetServerList.Server {
		queryArgs := make(map[string]interface{})
		queryArgs["server_uuid"] = server.UUID

		resGetVolumeList, err := queryparserext.GetVolumeList(queryArgs)
		if err == nil {
			for _, volume := range resGetVolumeList.(model.VolumeList).Volumes {
				volumesSize += volume.Size
			}
		}
	}

	if quota.LimitSSDGB+quota.LimitHDDGB < volumesSize {
		return errors.New("allocated volumes size is bigger than the quota limitation")
	}

	for _, pool := range resPoolList.Pool {
		if pool.Name == quota.PoolName {
			poolAvailableSize, _ := strconv.Atoi(pool.AvailableSize)
			if quota.LimitSSDGB+quota.LimitHDDGB-volumesSize > poolAvailableSize {
				return errors.New("not enough available size left for the pool")
			}

			break
		}
	}

	return nil
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

	if !isMaster && loginUserGroupID != groupID {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't create the other group's quota if you are not a master")}, nil
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

	err = checkPoolSize(quota, nil)
	if err != nil {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			err.Error())}, nil
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

func checkQuotaUpdate(quota model.Quota) error {
	var allocatedCPUCores = 0
	var allocatedMemoryGB = 0

	resGetServerList, err := client.RC.GetServerList(
		&pb.ReqGetServerList{
			Server: &pb.Server{
				GroupID: quota.GroupID,
			},
		})
	if err != nil {
		return errors.New("failed to get server list")
	}
	for _, server := range resGetServerList.Server {
		allocatedCPUCores += int(server.CPU)
		allocatedMemoryGB += int(server.Memory)
	}

	if quota.LimitCPUCores < allocatedCPUCores {
		return errors.New("allocated CPU cores to the group are bigger than the quota limitation")
	}
	if quota.LimitMemoryGB < allocatedMemoryGB {
		return errors.New("allocated memory to the group is bigger than the quota limitation")
	}

	resGetSubnetNum, err := client.RC.GetSubnetNum(
		&pb.ReqGetSubnetNum{
			GroupID: quota.GroupID,
		})
	if err != nil {
		return errors.New("failed to get subnets count")
	}

	if quota.LimitSubnetCnt < int(resGetSubnetNum.Num) {
		return errors.New("allocated subnets count is bigger than the quota limitation")
	}

	resGetAdaptiveIPServerNum, err := client.RC.GetAdaptiveIPServerNum(
		&pb.ReqGetAdaptiveIPServerNum{
			GroupID: quota.GroupID,
		})
	if err != nil {
		return errors.New("failed to get AdaptiveIPs count")
	}

	if quota.LimitSubnetCnt < int(resGetAdaptiveIPServerNum.Num) {
		return errors.New("allocated AdaptiveIPs count is bigger than the quota limitation")
	}

	err = checkPoolSize(quota, resGetServerList)
	if err != nil {
		return err
	}

	return nil
}

// UpdateQuota : Update the quota of the group
func UpdateQuota(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
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

	if !isMaster && loginUserGroupID != groupID {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group's quota if you are not a master")}, nil
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

	previousNodeList, err := client.RC.GetNodeList(
		&pb.ReqGetNodeList{
			Node: &pb.Node{
				GroupID: int64(groupID),
			},
		})
	if err != nil {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError,
			"Failed to get the previous quota info")}, nil
	}

	var duplicatedNodeUUIDs []string

	for _, nodeUUID := range selectedNodesSplit {
		var skipUpdate = false

		for _, previousNode := range previousNodeList.Node {
			if previousNode.UUID == nodeUUID {
				skipUpdate = true
				duplicatedNodeUUIDs = append(duplicatedNodeUUIDs, previousNode.UUID)
				break
			}
		}

		if skipUpdate {
			continue
		}

		_, err = client.RC.UpdateNode(&pb.ReqUpdateNode{
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

	for _, previousNode := range previousNodeList.Node {
		var duplicated = false

		for _, nodeUUID := range duplicatedNodeUUIDs {
			if nodeUUID == previousNode.UUID {
				duplicated = true
				break
			}
		}

		if duplicated {
			continue
		}

		_, err = client.RC.UpdateNode(&pb.ReqUpdateNode{
			Node: &pb.Node{
				UUID:    previousNode.UUID,
				GroupID: int64(-1),
			},
		})
		if err != nil {
			return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError,
				"Failed to delete groupID from the node (nodeUUID="+previousNode.UUID+")")}, nil
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

	err = checkQuotaUpdate(quota)
	if err != nil {
		return model.Quota{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			err.Error())}, nil
	}

	sql := "update quota set"
	sql += " limit_cpu_cores = " + strconv.Itoa(limitCPUCores) + ", "
	sql += " limit_memory_gb = " + strconv.Itoa(limitMemoryGB) + ", "
	sql += " limit_subnet_cnt = " + strconv.Itoa(subnetCnt) + ", "
	sql += " limit_adaptive_ip_cnt = " + strconv.Itoa(adaptiveCnt) + ", "
	sql += " pool_name = '" + poolName + "', "
	sql += " limit_ssd_gb = " + strconv.Itoa(ssdSize) + ", "
	sql += " limit_hdd_gb = " + strconv.Itoa(hddSize)
	sql += " where group_id = ?"

	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "UpdateQuota(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError,
			errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.Exec(groupID)
	if err != nil {
		errStr := "UpdateQuota(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.Node{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError,
			errStr)}, nil
	}

	return &quota, nil
}
