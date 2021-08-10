package mutationparser

import (
	"errors"
	"hcc/piccolo/action/graphql/queryparserext"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"math/bits"
	"strings"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// SignUp : Do user sign up process
func SignUp(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(int)
	id, idOk := args["id"].(string)
	authentication, authenticationOk := args["authentication"].(string)
	password, passwordOk := args["password"].(string)
	name, nameOk := args["name"].(string)
	email, emailOk := args["email"].(string)

	if !groupIDOk || !idOk || !authenticationOk || !passwordOk || !nameOk || !emailOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need id and authentication, group_id, password, name, email arguments")}, nil
	}

	if !isMaster && loginUserGroupID != groupID {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't create the other group's user if you are not a master")}, nil
	}

	if strings.ToLower(id) == "master" {
		logger.Logger.Println("SignUp(): Someone tried to sign up with master ID.")
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Hey, you can't be the master!")}, nil
	}

	authentication = strings.ToLower(authentication)
	if authentication != "admin" && authentication != "user" {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Wrong authentication provided!")}, nil
	}

	sql := "select id from user where id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &id)
	if err == nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Provided ID is in use")}, nil
	}

	_, err = dao.ReadGroup(groupID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "Provided Group ID is not exist")}, nil
		}

		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user := model.User{
		ID:             id,
		GroupID:        int64(groupID),
		Authentication: authentication,
		Name:           name,
		Email:          email,
	}

	sql = "insert into user(id, group_id, authentication, password, name, email, login_at, created_at) values (?, ?, ?, ?, ?, ?, now(), now())"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.ID, user.GroupID, authentication, password, user.Name, user.Email)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return &user, nil
}

// Unregister : Do user unregister process
func Unregister(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserID string, loginUserGroupID int) (interface{}, error) {
	id, idOk := args["id"].(string)

	if !idOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a id argument")}, nil
	}

	if strings.ToLower(id) == "master" {
		logger.Logger.Println("Unregister(): Someone tried to unregister master ID.")
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "You can't delete administrative IDs")}, nil
	}

	user, _ := queryparserext.User(args)
	if len(user.(model.User).Errors) != 0 && user.(model.User).Errors[0].ErrCode != 0 {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "user not found")}, nil
	}

	if !isMaster && int(user.(model.User).GroupID) != loginUserGroupID {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't unregister the other group's user if you are not a master")}, nil
	}

	if !isMaster && !isAdmin && user.(model.User).ID != loginUserID {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't unregister the other user if you are not a master or the admin")}, nil
	}

	sql := "delete from user where id = ?"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "Unregister(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		errStr := "Unregister(): " + err2.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	return model.User{ID: id, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}

// UpdateUser : Update info of the user
func UpdateUser(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	id, idOk := args["id"].(string)

	if !idOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a id argument")}, nil
	}

	authentication, authenticationOk := args["authentication"].(string)
	password, passwordOk := args["password"].(string)
	name, nameOk := args["name"].(string)
	email, emailOk := args["email"].(string)

	if !authenticationOk && !passwordOk && !nameOk && !emailOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need some arguments")}, nil
	}

	if !isMaster {
		user, err := queryparserext.User(args)
		if err != nil {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "Failed to get user info")}, nil
		}
		if loginUserGroupID != int(user.(model.User).GroupID) {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group's user if you are not a master")}, nil
		}
	}

	sql := "update user set"
	var updateSet = ""
	if authenticationOk {
		authentication = strings.ToLower(authentication)
		isWrongAuthentication := false
		if isMaster && id == "master" {
			isWrongAuthentication = authentication != "master"
		} else {
			isWrongAuthentication = authentication != "admin" && authentication != "user"
		}
		if isWrongAuthentication {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Wrong authentication provided!")}, nil
		}
		updateSet += " authentication = '" + authentication + "', "
	}
	if passwordOk {
		updateSet += " password = '" + password + "', "
	}
	if nameOk {
		updateSet += " name = '" + name + "', "
	}
	if emailOk {
		updateSet += " email = '" + email + "', "
	}

	sql += updateSet[0:len(updateSet)-2] + " where id = ?"

	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "UpdateUser(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err2 := stmt.Exec(id)
	if err2 != nil {
		errStr := "UpdateUser(): " + err2.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	_user, err := queryparserext.User(args)
	if err != nil {
		logger.Logger.Println("UpdateUser(): " + err.Error())
	}

	var user = _user.(model.User)

	user.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return &user, nil
}

func generateGroupID() (int64, error) {
	var groupID int64
	var groupIDmax int64

	sql := "select id from piccolo.group where id = 1000"
	row := mysql.Db.QueryRow(sql)
	err := mysql.QueryRowScan(row, &groupID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return 1000, nil
		}

		return 0, err
	}

	sql = "select max(id) from piccolo.group"
	row = mysql.Db.QueryRow(sql)
	err = mysql.QueryRowScan(row, &groupIDmax)
	if err != nil {
		return 0, err
	}

	if groupIDmax == (1<<bits.UintSize)/2-1 {
		return 0, errors.New("possible max group ID exceeded. You can also specify Group ID manually")
	}

	groupID = groupIDmax + 1

	return groupID, nil
}

func CreateGroup(args map[string]interface{}, isMaster bool) (interface{}, error) {
	if !isMaster {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	var groupID int64
	var err error

	_groupID, _groupIDOk := args["group_id"].(int)
	groupName, groupNameOk := args["group_name"].(string)

	if !_groupIDOk || !groupNameOk {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a group_name argument")}, nil
	}

	if _groupIDOk && _groupID != 0 {
		sql := "select id from piccolo.group where id = ?"
		row := mysql.Db.QueryRow(sql, _groupID)
		err := mysql.QueryRowScan(row, &_groupID)
		if err == nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Provided Group ID is in use")}, nil
		}

		if _groupID <= 0 {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Wrong Group ID provided should be bigger than 0")}, nil
		}

		groupID = int64(_groupID)
	} else {
		groupID, err = generateGroupID()
		if err != nil {
			return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
		}
	}

	group := model.Group{
		ID:   groupID,
		Name: groupName,
	}

	sql := "insert into piccolo.group(id, name) values (?, ?)"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(group.ID, group.Name)
	if err != nil {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	group.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return &group, nil
}

func UpdateGroup(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster && !isAdmin {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(int)
	groupName, groupNameOk := args["group_name"].(string)

	if !groupIDOk || !groupNameOk {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need group_id and group_name arguments")}, nil
	}

	if !isMaster && loginUserGroupID != groupID {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the other group if you are not a master")}, nil
	}

	if groupID == 1 {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't update the master group")}, nil
	}

	sql := "update piccolo.group set name = ? where id = ?"

	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "UpdateGroup(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err2 := stmt.Exec(groupName, groupID)
	if err2 != nil {
		errStr := "UpdateGroup(): " + err2.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	_group, err := dao.ReadGroup(groupID)
	if err != nil {
		logger.Logger.Println("UpdateGroup(): " + err.Error())
	}

	if _group != nil {
		_group.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return _group, nil
}

func DeleteGroup(args map[string]interface{}, isMaster bool) (interface{}, error) {
	if !isMaster {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(int)

	if !groupIDOk {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a group_id argument")}, nil
	}

	var userCount int
	sql := "select count(*) from piccolo.user where group_id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row, &userCount)
	if err != nil {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "Failed to get user count")}, nil
	}
	if userCount > 0 {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Provided Group ID is in use by some users")}, nil
	}

	group := model.Group{
		ID: int64(groupID),
	}

	sql = "delete from piccolo.group where id = ?"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(group.ID)
	if err != nil {
		return model.Group{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	group.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return &group, nil
}
