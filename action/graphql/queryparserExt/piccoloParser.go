package queryparserExt

import (
	dbsql "database/sql"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"time"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// User : Get the user info
func User(args map[string]interface{}) (interface{}, error) {
	var authentication string
	var name string
	var groupID int64
	var groupName string
	var email string
	var loginAt time.Time
	var createdAt time.Time

	uuid, uuidOk := args["uuid"].(string)
	id, idOk := args["id"].(string)

	sql := "select piccolo.user.uuid, piccolo.user.id, piccolo.user.authentication, piccolo.user.name, piccolo.user.group_id, piccolo.group.name as group_name, piccolo.user.email, piccolo.user.login_at, piccolo.user.created_at from piccolo.user, piccolo.group where piccolo.user.group_id = piccolo.group.id and"

	var row *dbsql.Row
	var err error

	if idOk && uuidOk {
		sql += " piccolo.user.id = ? and piccolo.user.uuid = ? order by piccolo.user.created_at"
		row = mysql.Db.QueryRow(sql, id, uuid)
	} else if idOk {
		sql += " piccolo.user.id = ? order by piccolo.user.created_at"
		row = mysql.Db.QueryRow(sql, id)
	} else if uuidOk {
		sql += " piccolo.user.uuid = ? order by piccolo.user.created_at"
		row = mysql.Db.QueryRow(sql, uuid)
	} else {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert uuid or id arguments")}, nil
	}

	err = mysql.QueryRowScan(row, &uuid, &id, &authentication, &name, &groupID, &groupName, &email, &loginAt, &createdAt)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user := model.User{UUID: uuid, ID: id, Authentication: authentication, Name: name,
		GroupID: groupID, GroupName: groupName,
		Email: email, LoginAt: loginAt, CreatedAt: createdAt}
	user.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return user, nil
}
