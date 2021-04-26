package queryparserext

import (
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

	id, idOk := args["id"].(string)

	if !idOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert uuid or id arguments")}, nil
	}

	sql := "select piccolo.user.id, piccolo.user.authentication, piccolo.user.name, piccolo.user.group_id, " +
		"piccolo.group.name as group_name, piccolo.user.email, piccolo.user.login_at, piccolo.user.created_at from " +
		"piccolo.user, piccolo.group where piccolo.user.group_id = piccolo.group.id and piccolo.user.id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &id, &authentication, &name, &groupID, &groupName, &email, &loginAt, &createdAt)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user := model.User{ID: id, Authentication: authentication, Name: name,
		GroupID: groupID, GroupName: groupName,
		Email: email, LoginAt: loginAt, CreatedAt: createdAt}
	user.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return user, nil
}
