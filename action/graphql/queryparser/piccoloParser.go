package queryparser

import (
	dbsql "database/sql"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
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
	var uuid string
	var loginAt time.Time
	var createdAt time.Time
	var noLimit bool

	id, idOk := args["id"].(string)
	authentication, authenticationOk := args["authentication"].(string)
	name, nameOk := args["name"].(string)
	groupID, groupIDOk := args["group_id"].(int)
	groupName, groupNameOk := args["group_name"].(string)
	email, emailOk := args["email"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		noLimit = true
	} else if rowOk && pageOk {
		noLimit = false
	} else {
		return model.UserList{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	sqlSelect := "select piccolo.user.uuid, piccolo.user.id, piccolo.user.authentication, piccolo.user.name, piccolo.user.group_id, piccolo.group.name as group_name, piccolo.user.email, piccolo.user.login_at, piccolo.user.created_at"
	sqlCount := "select count(*)"
	sql := " from piccolo.user, piccolo.group where piccolo.user.group_id = piccolo.group.id"

	if idOk {
		sql += " and piccolo.user.id = '" + id + "'"
	}
	if authenticationOk {
		sql += " and piccolo.user.authentication = '" + authentication + "'"
	}
	if nameOk {
		sql += " and piccolo.user.name = '" + name + "'"
	}
	if groupIDOk {
		sql += " and piccolo.user.group_id = " + strconv.Itoa(groupID)
	}
	if groupNameOk {
		sql += " and group_name = '" + groupName + "'"
	}
	if emailOk {
		sql += " and piccolo.user.email = '" + email + "'"
	}

	if !noLimit {
		sql += " order by piccolo.user.created_at desc limit ? offset ?"
	}

	var stmt *dbsql.Rows
	var err error

	var userNum int
	result := mysql.Db.QueryRow(sqlCount + sql)
	err = mysql.QueryRowScan(result, &userNum)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	if noLimit {
		stmt, err = mysql.Query(sqlSelect + sql)
	} else {
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
		err := stmt.Scan(&uuid, &id, &authentication, &name, &groupID, &groupName, &email, &loginAt, &createdAt)
		if err != nil {
			logger.Logger.Println(err)
		}
		user := model.User{UUID: uuid, ID: id, Authentication: authentication, Name: name,
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

// CheckToken : Do token validation check process
func CheckToken(args map[string]interface{}) (interface{}, error) {
	_, tokenOk := args["token"].(string)
	if !tokenOk {
		return model.IsValid{IsValid: false, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a token argument")}, nil
	}

	_, _, _, _, err := usertool.ValidateToken(args, false)
	if err != nil {
		return model.IsValid{IsValid: false, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
	}

	return model.IsValid{IsValid: true, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
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
