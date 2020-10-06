package queryparser

import (
	dbsql "database/sql"
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
	"time"
)

func updateUserLoginAt(id string) error {
	sql := "update user set login_at = now() where id = ?"

	stmt, err := mysql.Db.Prepare(sql)
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
		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need id and password arguments")}, nil
	}

	var dbPassword string

	sql := "select password from user where id = ?"
	err := mysql.Db.QueryRow(sql, id).Scan(&dbPassword)
	if err != nil {
		logger.Logger.Println(err)

		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLLoginFailed, "user not found or password mismatch")}, nil
	}

	// Given password is hashed password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
	if err != nil {
		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLLoginFailed, "user not found or password mismatch")}, nil
	}

	logger.Logger.Println("User logged in: " + id)

	err = updateUserLoginAt(id)
	if err != nil {
		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	token, err := usertool.GenerateToken(id, password)
	if err != nil {
		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTokenGenerationError, err.Error())}, nil
	}

	return model.Token{Token: token, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
}

// User : Get the user info
func User(args map[string]interface{}) (interface{}, error) {
	var name string
	var email string
	var loginAt time.Time
	var createdAt time.Time

	uuid, uuidOk := args["uuid"].(string)
	id, idOk := args["id"].(string)

	sql := "select uuid, id, name, email, login_at, created_at from user where"

	var row *dbsql.Row
	var err error

	if idOk && uuidOk {
		sql += " id = ? and uuid = ? order by created_at"
		row = mysql.Db.QueryRow(sql, id, uuid)
	} else if idOk {
		sql += " id = ? order by created_at"
		row = mysql.Db.QueryRow(sql, id)
	} else if uuidOk {
		sql += " uuid = ? order by created_at"
		row = mysql.Db.QueryRow(sql, uuid)
	} else {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "please insert uuid or id arguments")}, nil
	}

	err = row.Scan(&uuid, &id, &name, &email, &loginAt, &createdAt)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user := model.User{UUID: uuid, ID: id, Name: name, Email: email, LoginAt: loginAt, CreatedAt: createdAt}
	user.Errors = errors.ReturnHccEmptyErrorPiccolo()

	return user, nil
}

// UserList : Get the user list
func UserList(args map[string]interface{}) (interface{}, error) {
	var users []model.User
	var uuid string
	var loginAt time.Time
	var createdAt time.Time
	var noLimit bool

	id, idOk := args["id"].(string)
	name, nameOk := args["name"].(string)
	email, emailOk := args["email"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		noLimit = true
	} else if rowOk && pageOk {
		noLimit = false
	} else {
		return model.UserList{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	sql := "select uuid, id, name, email, login_at, created_at from user where 1=1"

	if idOk {
		sql += " and id = '" + id + "'"
	}
	if nameOk {
		sql += " and name = '" + name + "'"
	}
	if emailOk {
		sql += " and email = '" + email + "'"
	}

	if !noLimit {
		sql += " order by created_at desc limit ? offset ?"
	}

	var stmt *dbsql.Rows
	var err error

	if noLimit {
		stmt, err = mysql.Db.Query(sql)
	} else {
		stmt, err = mysql.Db.Query(sql, row, row*(page-1))
	}

	if err != nil {
		logger.Logger.Println(err)
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err := stmt.Scan(&uuid, &id, &name, &email, &loginAt, &createdAt)
		if err != nil {
			logger.Logger.Println(err)
		}
		user := model.User{UUID: uuid, ID: id, Name: name, Email: email, LoginAt: loginAt, CreatedAt: createdAt}
		users = append(users, user)
	}

	return model.UserList{Users: users, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
}

// CheckToken : Do token validation check process
func CheckToken(args map[string]interface{}) (interface{}, error) {
	_, tokenOk := args["token"].(string)
	if !tokenOk {
		return model.IsValid{IsValid: false, Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a token argument")}, nil
	}

	err := usertool.ValidateToken(args)
	if err != nil {
		return model.IsValid{IsValid: false, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
	}

	return model.IsValid{IsValid: true, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
}

// ResourceUsage : Get usage of resources
func ResourceUsage() (interface{}, error) {
	resGetNodeList, err := client.RC.GetNodeList(&rpcflute.ReqGetNodeList{Node: &rpcflute.Node{}})
	if err != nil {
		return model.ResourceUsage{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, "failed to get nodes")}, nil
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

	// TODO : Currently, total storage is hard coded.
	// FIXME : Need to fix to get total storage from cello module.
	total.Storage = 2048

	for _, node := range resGetNodeList.Node {
		if node.Active == 1 {
			inUse.CPU += int(node.CPUCores)
			inUse.Memory += int(node.Memory)
			// TODO : Currently, in-use storage is hard coded.
			// FIXME : Need to fix to get in-use storage from cello module.
			inUse.Storage += 10
			inUse.Node++
		}

		total.CPU += int(node.CPUCores)
		total.Memory += int(node.Memory)
		total.Node++
	}

	return model.ResourceUsage{Total: total, InUse: inUse, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
}
