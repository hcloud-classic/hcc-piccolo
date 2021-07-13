package queryparser

import (
	dbsql "database/sql"
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/graphql/queryparserext"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
	"strconv"
	"strings"
	"time"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"

	"golang.org/x/crypto/bcrypt"
)

func updateUserLoginAt(id string) error {
	sql := "update user set login_at = now() where id = ?"

	stmt, err := mysql.Prepare(sql)
	if err != nil {
		logger.Logger.Println("updateUserLoginAt(): " + err.Error())
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err2 := stmt.Exec(id)
	if err2 != nil {
		logger.Logger.Println("updateUserLoginAt(): " + err2.Error())
		return err2
	}

	return nil
}

// Login : Do user login process
func Login(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)
	password, passwordOk := args["password"].(string)

	if !idOk || !passwordOk {
		return model.Token{Token: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need id and password arguments")}, nil
	}

	var dbPassword string

	sql := "select password from user where id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &dbPassword)
	if err != nil {
		logger.Logger.Println(err)

		return model.Token{Token: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLLoginFailed, "user not found or password mismatch")}, nil
	}

	// Given password is hashed password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
	if err != nil {
		return model.Token{Token: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLLoginFailed, "user not found or password mismatch")}, nil
	}

	if strings.ToLower(id) == "admin" || strings.ToLower(id) == "administrator" {
		logger.Logger.Println("ADMIN LOGGED IN")
	} else {
		logger.Logger.Println("User logged in: " + id)
	}

	err = updateUserLoginAt(id)
	if err != nil {
		return model.Token{Token: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	token, err := usertool.GenerateToken(id, password)
	if err != nil {
		return model.Token{Token: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLTokenGenerationError, err.Error())}, nil
	}

	return model.Token{Token: token, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// UserList : Get the user list
func UserList(args map[string]interface{}) (interface{}, error) {
	var users []model.User
	var loginAt time.Time
	var createdAt time.Time
	var noLimit bool

	id, idOk := args["id"].(string)
	authentication, authenticationOk := args["authentication"].(string)
	name, nameOk := args["name"].(string)
	groupID, groupIDOk := args["group_id"].(int)
	if groupID == 0 {
		groupIDOk = false
	}
	groupName, groupNameOk := args["group_name"].(string)
	email, emailOk := args["email"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		noLimit = true
	} else if rowOk && pageOk {
		noLimit = false
		if row == 0 && page == 0 {
			noLimit = true
		}
	} else {
		return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	sqlSelect := "select piccolo.user.id, piccolo.user.authentication, piccolo.user.name, piccolo.user.group_id, piccolo.group.name as group_name, piccolo.user.email, piccolo.user.login_at, piccolo.user.created_at"
	sqlCount := "select count(*)"
	sql := " from piccolo.user, piccolo.group where piccolo.user.group_id = piccolo.group.id"

	if idOk && len(id) != 0 {
		sql += " and piccolo.user.id like '%" + id + "%'"
	}
	if authenticationOk && len(authentication) != 0 {
		sql += " and piccolo.user.authentication like '%" + authentication + "%'"
	}
	if nameOk && len(name) != 0 {
		sql += " and piccolo.user.name like '%" + name + "%'"
	}
	if groupIDOk && groupID != 0 {
		sql += " and piccolo.user.group_id = " + strconv.Itoa(groupID)
	}
	if groupNameOk && len(groupName) != 0 {
		sql += " and piccolo.group.name like '%" + groupName + "%'"
	}
	if emailOk && len(email) != 0 {
		sql += " and piccolo.user.email like '%" + email + "%'"
	}

	var stmt *dbsql.Rows
	var err error

	var userNum int
	result := mysql.Db.QueryRow(sqlCount + sql)
	err = mysql.QueryRowScan(result, &userNum)
	if err != nil {
		return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	if noLimit {
		stmt, err = mysql.Query(sqlSelect + sql)
	} else {
		sql += " order by piccolo.user.created_at desc limit ? offset ?"
		stmt, err = mysql.Query(sqlSelect+sql, row, row*(page-1))
	}

	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err := stmt.Scan(&id, &authentication, &name, &groupID, &groupName, &email, &loginAt, &createdAt)
		if err != nil {
			logger.Logger.Println(err)
		}
		user := model.User{ID: id, Authentication: authentication, Name: name,
			GroupID: int64(groupID), GroupName: groupName,
			Email: email, LoginAt: loginAt, CreatedAt: createdAt}
		users = append(users, user)
	}

	return model.UserList{Users: users, TotalNum: userNum, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// NumUser : Get number of users
func NumUser(args map[string]interface{}) (interface{}, error) {
	var userNum int

	groupID, groupIDOk := args["group_id"].(int)

	sql := "select count(*) from user"
	if groupIDOk {
		sql = "select count(*) from user where group_id = " + strconv.Itoa(groupID)
	}
	row := mysql.Db.QueryRow(sql)
	err := mysql.QueryRowScan(row, &userNum)
	if err != nil {
		return model.UserNum{Number: userNum, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	return model.UserNum{Number: userNum, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// ReadGroupList : Get list of groups
func ReadGroupList(args map[string]interface{}, isMaster bool) (interface{}, error) {
	includeMaster, includeMasterOk := args["include_master"].(bool)

	if !isMaster {
		return model.GroupList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	var groupList []model.Group

	resGetGroupList, errCode, errText := dao.ReadGroupList()
	if errCode != 0 || errText != "" {
		return model.GroupList{Errors: errconv.ReturnHccErrorPiccolo(errCode, errText)}, nil
	}

	for i := range resGetGroupList.Group {
		group := pbtomodel.PbGroupToModelGroup(resGetGroupList.Group[i])
		if (!includeMasterOk || (includeMasterOk && !includeMaster)) &&
			group.ID == 1 {
			continue
		}
		groupList = append(groupList, *group)
	}

	return model.GroupList{Groups: groupList}, nil
}

// CheckToken : Do token validation check process
func CheckToken(args map[string]interface{}) (interface{}, error) {
	_, tokenOk := args["token"].(string)
	if !tokenOk {
		return model.IsValid{IsValid: false, Authentication: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a token argument")}, nil
	}

	isAdmin, isMaster, id, groupID, err := usertool.ValidateToken(args, false)
	if err != nil {
		return model.IsValid{IsValid: false, Authentication: "", Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
	}

	var authentication = "user"
	if isAdmin {
		authentication = "admin"
	} else if isMaster {
		authentication = "master"
	}

	group, err := dao.ReadGroup(int(groupID))
	if err != nil {
		return model.IsValid{IsValid: false, Authentication: "", Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	return model.IsValid{
		IsValid:        true,
		UserID:         id,
		GroupID:        group.ID,
		GroupName:      group.Name,
		Authentication: authentication,
		Errors:         errconv.ReturnHccEmptyErrorPiccolo(),
	}, nil
}

// ResourceUsage : Get usage of resources
func ResourceUsage(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)

	var reqGetNodeList pb.ReqGetNodeList
	reqGetNodeList.Node = &pb.Node{}
	if groupIDOk {
		reqGetNodeList.Node.GroupID = int64(groupID)
	}

	resGetNodeList, err := client.RC.GetNodeList(&reqGetNodeList)
	if err != nil {
		return model.ResourceUsage{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "failed to get nodes")}, nil
	}

	var total model.Resource
	var inUse model.Resource

	total.CPU = 0
	total.Memory = 0
	total.Storage = 0
	total.Node = 0
	inUse.CPU = 0
	inUse.Memory = 0
	inUse.Storage = 0
	inUse.Node = 0

	poolArg := map[string]interface{}{
		"action":        "read",
		"uuid":          "",
		"size":          "",
		"free":          "",
		"capacity":      "",
		"health":        "",
		"name":          "",
		"availablesize": "",
	}

	// TODO: Need to handle group_id - ish
	poolStruct, err := GetPoolList(poolArg)
	convModelPools := poolStruct.(model.PoolList)
	for _, eachPool := range convModelPools.Pools {
		tempSize, _ := strconv.Atoi(eachPool.Size)
		total.Storage += tempSize
		tempUsed, _ := strconv.Atoi(eachPool.Used)
		inUse.Storage += tempUsed
	}

	for _, node := range resGetNodeList.Node {
		if node.Active == 1 {
			inUse.CPU += int(node.CPUCores)
			inUse.Memory += int(node.Memory)

			inUse.Node++
		}
		total.CPU += int(node.CPUCores)
		total.Memory += int(node.Memory)
		total.Node++
	}

	return model.ResourceUsage{Total: total, InUse: inUse, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// QuotaList : Get the quota list
func QuotaList(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.QuotaList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	var quotas []model.Quota
	var noLimit bool

	groupID, groupIDOk := args["group_id"].(int)
	if groupID == 0 {
		groupIDOk = false
	}
	groupName, groupNameOk := args["group_name"].(string)

	if !isMaster && (groupIDOk || groupNameOk) {
		return model.QuotaList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Only a master can search other group's quotas!")}, nil
	}

	limitCPUCores, limitCPUCoresOk := args["limit_cpu_cores"].(int)
	limitMemoryGB, limitMemoryGBOk := args["limit_memory_gb"].(int)
	limitSubnetCnt, limitSubnetCntOk := args["limit_subnet_cnt"].(int)
	limitAdaptiveIPCnt, limitAdaptiveIPCntOk := args["limit_adaptive_ip_cnt"].(int)
	poolName, poolNameOk := args["pool_name"].(string)
	limitSSDGB, limitSSDGBOk := args["limit_ssd_gb"].(int)
	limitHDDGB, limitHDDGBOk := args["limit_hdd_gb"].(int)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		noLimit = true
	} else if rowOk && pageOk {
		noLimit = false
		if row == 0 && page == 0 {
			noLimit = true
		}
	} else {
		return model.QuotaList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	sqlSelect := "select piccolo.quota.group_id, piccolo.group.name as group_name, " +
		"piccolo.quota.limit_cpu_cores, piccolo.quota.limit_memory_gb, " +
		"piccolo.quota.limit_subnet_cnt, piccolo.quota.limit_adaptive_ip_cnt, " +
		"piccolo.quota.pool_name, piccolo.quota.limit_ssd_gb, piccolo.quota.limit_hdd_gb"
	sqlCount := "select count(*)"
	sql := " from piccolo.quota, piccolo.group where piccolo.quota.group_id = piccolo.group.id"

	if !isMaster {
		sql = " from piccolo.quota, piccolo.group where piccolo.quota.group_id = piccolo.group.id and " +
			"piccolo.group.id = " + strconv.Itoa(loginUserGroupID)
	}

	if groupIDOk && groupID != 0 {
		sql += " and piccolo.quota.group_id = " + strconv.Itoa(groupID)
	}
	if groupNameOk && len(groupName) != 0 {
		sql += " and piccolo.group.name like '%" + groupName + "%'"
	}
	if limitCPUCoresOk && limitCPUCores != 0 {
		sql += " and piccolo.quota.limit_cpu_cores = " + strconv.Itoa(limitCPUCores)
	}
	if limitMemoryGBOk && limitMemoryGB != 0 {
		sql += " and piccolo.quota.limit_memory_gb = " + strconv.Itoa(limitMemoryGB)
	}
	if limitSubnetCntOk && limitSubnetCnt != 0 {
		sql += " and piccolo.quota.limit_subnet_cnt = " + strconv.Itoa(limitSubnetCnt)
	}
	if limitAdaptiveIPCntOk && limitAdaptiveIPCnt != 0 {
		sql += " and piccolo.quota.limit_adaptive_ip_cnt = " + strconv.Itoa(limitAdaptiveIPCnt)
	}
	if poolNameOk && len(poolName) != 0 {
		sql += " and piccolo.quota.pool_name like '%" + poolName + "%'"
	}
	if limitSSDGBOk && limitSSDGB != 0 {
		sql += " and piccolo.quota.limit_ssd_gb = " + strconv.Itoa(limitSSDGB)
	}
	if limitHDDGBOk && limitHDDGB != 0 {
		sql += " and piccolo.quota.limit_hdd_gb = " + strconv.Itoa(limitHDDGB)
	}

	var stmt *dbsql.Rows
	var err error

	var quotaNum int
	result := mysql.Db.QueryRow(sqlCount + sql)
	err = mysql.QueryRowScan(result, &quotaNum)
	if err != nil {
		return model.QuotaList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	if noLimit {
		stmt, err = mysql.Query(sqlSelect + sql)
	} else {
		sql += " order by piccolo.quota.group_id asc limit ? offset ?"
		stmt, err = mysql.Query(sqlSelect+sql, row, row*(page-1))
	}

	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err := stmt.Scan(&groupID, &groupName, &limitCPUCores, &limitMemoryGB, &limitSubnetCnt, &limitAdaptiveIPCnt, &poolName, &limitSSDGB, &limitHDDGB)
		if err != nil {
			logger.Logger.Println(err)
		}
		quota := model.Quota{
			GroupID:            int64(groupID),
			GroupName:          groupName,
			LimitCPUCores:      limitCPUCores,
			LimitMemoryGB:      limitMemoryGB,
			LimitSubnetCnt:     limitSubnetCnt,
			LimitAdaptiveIPCnt: limitAdaptiveIPCnt,
			PoolName:           poolName,
			LimitSSDGB:         limitSSDGB,
			LimitHDDGB:         limitHDDGB,
		}
		quotas = append(quotas, quota)
	}

	return model.QuotaList{Quotas: quotas, TotalNum: quotaNum, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

func getGroup(groupID int) (*model.Group, error) {
	var group model.Group

	sql := "select piccolo.group.id, piccolo.group.name from piccolo.group where id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row, &group.ID, &group.Name)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// QuotaDetail : Get detail info of the quota
func QuotaDetail(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(int)
	if !groupIDOk {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a group_id argument")}, nil
	}

	var quataDetail model.QuotaDetail

	group, err := getGroup(groupID)
	if err != nil {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "Failed to get group info!")}, nil
	}

	// Group
	quataDetail.GroupID = group.ID
	quataDetail.GroupName = group.Name

	queryArgs := make(map[string]interface{})
	queryArgs["group_id"] = int(group.ID)

	// Nodes
	var totalCPUCores = 0
	var totalMemoryGB = 0
	nodes, err := ListNode(queryArgs)
	if err != nil {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "Failed to get info of nodes!")}, nil
	}
	serverUUIDs := make(map[string]string)
	quataDetail.Nodes = nodes.(model.NodeList).Nodes
	for _, node := range quataDetail.Nodes {
		totalCPUCores += node.CPUCores
		totalMemoryGB += node.Memory
		if node.Active == 1 {
			serverUUIDs[node.ServerUUID] = node.ServerUUID
		}
	}
	quataDetail.TotalCPUCores = totalCPUCores
	quataDetail.TotalMemoryGB = totalMemoryGB

	// TotalSubnetNum
	subnets, err := ListSubnet(queryArgs)
	if err != nil {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "Failed to get info of subnets!")}, nil
	}
	quataDetail.TotalSubnetNum = len(subnets.(model.SubnetList).Subnets)

	// TotalAdaptiveIPNum
	adaptiveIPNum, err := NumAdaptiveIPServer(queryArgs)
	if err != nil {
		return model.QuotaDetail{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, "Failed to get number of AdaptiveIPs!")}, nil
	}
	quataDetail.TotalAdaptiveIPNum = adaptiveIPNum.(model.AdaptiveIPServerNum).Number

	// Volumes
	var volumes []model.Volume
	for _, serverUUID := range serverUUIDs {
		queryArgs := make(map[string]interface{})
		queryArgs["server_uuid"] = serverUUID

		volumeList, err := queryparserext.GetVolumeList(queryArgs)
		if err == nil {
			volumes = append(volumes, volumeList.(model.VolumeList).Volumes...)
		}
	}
	quataDetail.Volumes = volumes

	// Errors
	quataDetail.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return quataDetail, nil
}
