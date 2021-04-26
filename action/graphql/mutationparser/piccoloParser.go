package mutationparser

import (
	"hcc/piccolo/action/graphql/queryparserext"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"strings"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// SignUp : Do user sign up process
func SignUp(args map[string]interface{}, isAdmin bool, isMaster bool, loginUserGroupID int) (interface{}, error) {
	if !isMaster || !isAdmin {
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
