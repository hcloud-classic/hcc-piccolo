package dao

import (
	"errors"
	"hcc/piccolo/action/graphql/queryparserext"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"strconv"
	"strings"
	"time"

	"innogrid.com/hcloud-classic/hcc_errors"
)

func isAutoScaleTriggered(serverUUID string) (bool, string) {
	var detail string

	sql := "select detail from piccolo.server_alarm where server_uuid = ? and auto_scale_triggered = 1"
	row := mysql.Db.QueryRow(sql, serverUUID)
	err := mysql.QueryRowScan(row, &detail)
	if err != nil {
		return false, ""
	}

	return true, detail
}

func turnOffAutoScaleTriggered(serverUUID string) error {
	sql := "update piccolo.server_alarm set auto_scale_triggered = 0 where server_uuid = ?"
	row := mysql.Db.QueryRow(sql, serverUUID)
	err := mysql.QueryRowScan(row)
	if err != nil {
		return err
	}

	return nil
}

// WriteServerAlarm : Write a server alarm to the database
func WriteServerAlarm(serverUUID string, reason string, detail string) error {
	var autoScaleTriggered = 0

	if strings.Contains(reason, "AutoScale Queued") {
		_ = turnOffAutoScaleTriggered(serverUUID)
	}

	if strings.Contains(reason, "AutoScale Triggered") {
		autoScaleTriggered = 1
	}

	if autoScaleTriggered == 1 {
		triggered, detailStr := isAutoScaleTriggered(serverUUID)
		if triggered {
			if detail == detailStr {
				return nil
			}
			_ = turnOffAutoScaleTriggered(serverUUID)
		}
	}

	if autoScaleTriggered == 1 && detail == "Turn Off AutoScale" {
		err := turnOffAutoScaleTriggered(serverUUID)
		if err != nil {
			return err
		}
		reason = "AutoScale Trigger canceled"
		detail = "Server is back to normal. (ServerUUID: " + serverUUID + ")"
	}

	stmt, err := mysql.Prepare("insert into piccolo.server_alarm(user_id, server_uuid, reason, detail, time, auto_scale_triggered) values(?, ?, ?, ?, now(), ?)")
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	server, err := client.RC.GetServer(serverUUID)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(server.Server.UserUUID, serverUUID, reason, detail, autoScaleTriggered)
	if err != nil {
		return err
	}

	return nil
}

// DeleteServerAlarm : Delete alarms of the server from the database
func DeleteServerAlarm(args map[string]interface{}, loginUserID string) (interface{}, error) {
	no, noOk := args["no"].(int)
	if !noOk {
		return model.ServerAlarm{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a no argument")}, nil
	}

	var userID string

	sql := "select user_id from piccolo.server_alarm where no = ?"
	row := mysql.Db.QueryRow(sql, no)
	err := mysql.QueryRowScan(row, &userID)
	if err != nil {
		return model.ServerAlarm{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	if userID != loginUserID {
		return model.ServerAlarm{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "You can't delete other user's server alarms")}, nil
	}

	stmt, err := mysql.Prepare("delete from server_alarm where no = ?")
	if err != nil {
		errStr := "DeleteServerAlarm(): " + err.Error()
		logger.Logger.Println(errStr)

		return model.ServerAlarm{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(no)
	if err != nil {
		errStr := "DeleteServerAlarm(): " + err.Error()
		logger.Logger.Println(errStr)

		return model.ServerAlarm{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	return model.ServerAlarm{No: no}, nil
}

func getUserName(userID string) string {
	queryArgs := make(map[string]interface{})
	queryArgs["id"] = userID
	data, _ := queryparserext.User(queryArgs)

	return data.(model.User).Name
}

func unsetUnread(no int) error {
	sql := "update piccolo.server_alarm set unread = 0 where no = ?"

	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "updateUserAlarmTriggered(): " + err.Error()
		logger.Logger.Println(errStr)

		return errors.New(errStr)
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err2 := stmt.Exec(no)
	if err2 != nil {
		errStr := "unsetUnread(): " + err2.Error()
		logger.Logger.Println(errStr)

		return errors.New(errStr)
	}

	return nil
}

// ShowServerAlarms : Show server alarms from the database
func ShowServerAlarms(args map[string]interface{}) (interface{}, error) {
	var err error
	var alarms []model.ServerAlarm
	var serverAlarms model.ServerAlarms
	var totalNum = 0

	userID, userIDOk := args["user_id"].(string)
	if !userIDOk {
		return model.ServerAlarms{ServerAlarms: alarms, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a user_id argument")}, nil
	}

	var isLimit bool
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk && !pageOk {
		isLimit = false
	} else if rowOk && pageOk {
		isLimit = true
	} else {
		return model.ServerAlarms{ServerAlarms: alarms, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "please insert row and page arguments or leave arguments as empty state")}, nil
	}

	var no int
	var serverUUID string
	var reason string
	var detail string
	var _time time.Time
	var unread int
	var autoScaleTriggered int

	sql := "select no, server_uuid, reason, detail, time, unread, auto_scale_triggered from piccolo.server_alarm where user_id = ? order by no desc"
	if isLimit {
		sql += " limit " + strconv.Itoa(row) + " offset " + strconv.Itoa(row*(page-1))
	}

	stmt, err := mysql.Db.Query(sql, userID)
	if err != nil {
		goto ERROR
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err = stmt.Scan(&no, &serverUUID, &reason, &detail, &_time, &unread, &autoScaleTriggered)
		if err != nil {
			goto ERROR
		}

		resGetServer, err := client.RC.GetServer(serverUUID)
		if err != nil {
			goto ERROR
		}

		alarms = append(alarms, model.ServerAlarm{
			No:                 no,
			UserID:             userID,
			UserName:           getUserName(userID),
			ServerUUID:         serverUUID,
			ServerName:         resGetServer.Server.ServerName,
			Reason:             reason,
			Detail:             detail,
			Time:               _time,
			Unread:             unread,
			AutoScaleTriggered: autoScaleTriggered,
		})
		err = unsetUnread(no)
		if err != nil {
			goto ERROR
		}
		totalNum++
	}
	serverAlarms.TotalNum = totalNum

ERROR:
	serverAlarms.ServerAlarms = alarms
	if err != nil {
		serverAlarms.Errors = errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())
	} else {
		serverAlarms.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return serverAlarms, nil
}

// ShowUnreadServerAlarmsNum : Show number of unread server alarms from the database
func ShowUnreadServerAlarmsNum(args map[string]interface{}) (interface{}, error) {
	var err error
	var serverAlarmsNum model.ServerAlarmsNum
	var serverAlarmsNr int64

	serverAlarmsNum.Number = 0

	userID, userIDOk := args["user_id"].(string)
	if !userIDOk {
		return model.ServerAlarmsNum{Number: 0, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a user_id argument")}, nil
	}

	sql := "select count(*) from piccolo.server_alarm where user_id = ? and unread = 1"
	row := mysql.Db.QueryRow(sql, userID)
	err = mysql.QueryRowScan(row, &serverAlarmsNr)
	if err != nil {
		goto ERROR
	}

	serverAlarmsNum.Number = int(serverAlarmsNr)

ERROR:
	if err != nil {
		serverAlarmsNum.Errors = errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloInternalInitFail, err.Error())
	} else {
		serverAlarmsNum.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return serverAlarmsNum, nil
}
