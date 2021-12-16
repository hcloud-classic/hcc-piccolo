package dao

import (
	db "database/sql"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
	"strconv"
	"time"

	"innogrid.com/hcloud-classic/hcc_errors"
)

// WriteServerAction : Write a server alarm to the database
func WriteServerAction(serverUUID string, action string, result string, errStr string, token string) error {
	stmt, err := mysql.Prepare("insert into server_actions(server_uuid, action, result, err_str, user_id, time) values(?, ?, ?, ?, ?, now())")
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	userID, err := usertool.GetUserID(token)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(serverUUID, action, result, errStr, userID)
	if err != nil {
		return err
	}

	return nil
}

// *** We are using ARCHIVE engine for server_actions table ***
// DeleteServerAction : Delete logs of the server from the database
//func DeleteServerAction(serverUUID string) error {
//	stmt, err := mysql.Prepare("delete from server_actions where server_uuid = ?")
//	if err != nil {
//		return err
//	}
//
//	_, err = stmt.Exec(serverUUID)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// ShowServerActions : Show server actions from the database
func ShowServerActions(args map[string]interface{}) (interface{}, error) {
	var err error
	var actions []model.ServerAction
	var serverActions model.ServerActions

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return model.ServerActions{ServerActions: actions, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	var isLimit bool
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		isLimit = false
	} else if rowOk && pageOk {
		isLimit = true
	} else {
		return model.ServerActions{ServerActions: actions, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	var action string
	var result string
	var errStr string
	var userID string
	var _time time.Time
	var serverActionsNr int64

	var queryRow *db.Row

	sql := "select action, result, err_str, user_id, time from server_actions where server_uuid = ? order by no desc"
	if isLimit {
		sql += " limit " + strconv.Itoa(row) + " offset " + strconv.Itoa(row*(page-1))
	}

	stmt, err := mysql.Db.Query(sql, serverUUID)
	if err != nil {
		goto ERROR
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err = stmt.Scan(&action, &result, &errStr, &userID, &_time)
		if err != nil {
			goto ERROR
		}

		actions = append(actions, model.ServerAction{
			Action: action,
			Result: result,
			ErrStr: errStr,
			UserID: userID,
			Time:   _time,
		})
	}

	serverActions.Number = 0
	sql = "select count(*) from server_actions where server_uuid = ?"
	queryRow = mysql.Db.QueryRow(sql, serverUUID)
	err = mysql.QueryRowScan(queryRow, &serverActionsNr)
	if err != nil {
		goto ERROR
	}

	serverActions.Number = int(serverActionsNr)

ERROR:
	serverActions.ServerActions = actions
	if err != nil {
		serverActions.Errors = errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())
	} else {
		serverActions.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return serverActions, nil
}
