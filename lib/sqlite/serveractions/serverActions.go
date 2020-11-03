package serveractions

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3" // Needed for use sqlite3
	hccerr "hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/usertool"
	"os"
	"time"
)

func dbPath(serverUUID string) string {
	return "/var/log/" + logger.LogName + "/server_actions/"
}

func dbFile(serverUUID string) string {
	return dbPath(serverUUID) + "/" + serverUUID + ".db"
}

func createDatabase(serverUUID string) error {
	err := logger.CreateDirIfNotExist(dbPath(serverUUID))
	if err != nil {
		return err
	}

	_, err = os.OpenFile(dbFile(serverUUID), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", dbFile(serverUUID))
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE `server_actions` (`action` VARCHAR(255) NOT NULL, `result` VARCHAR(20) NOT NULL, `err_str` VARCHAR(255) NOT NULL, `user_id` VARCHAR(255) NOT NULL, `time` DATETIME NOT NULL)")
	if err != nil {
		return err
	}

	_ = db.Close()

	return nil
}

// WriteServerAction : Write a server action to the sqlite db file
func WriteServerAction(serverUUID string, action string, result string, errStr string, token string) error {
	if _, err := os.Stat(dbFile(serverUUID)); os.IsNotExist(err) {
		err = createDatabase(serverUUID)
		if err != nil {
			return err
		}
	}

	db, err := sql.Open("sqlite3", dbFile(serverUUID))
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	stmt, err := db.Prepare("INSERT INTO server_actions(action, result, err_str, user_id, time) values(?, ?, ?, ?, CURRENT_TIMESTAMP)")
	if err != nil {
		return err
	}

	userID, err := usertool.GetUserID(token)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(action, result, errStr, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteServerAction : Delete the sqlite db file of the server
func DeleteServerAction(serverUUID string) error {
	err := os.Remove(dbFile(serverUUID))
	if err != nil {
		return err
	}

	return nil
}

// ShowServerActions : Show server actions from the sqlite db file
func ShowServerActions(args map[string]interface{}) (interface{}, error) {
	var err error
	var actions []ServerAction
	var serverActions ServerActions

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return ServerActions{ServerActions: actions, Errors: hccerr.ReturnHccErrorPiccolo(hccerr.PiccoloGraphQLArgumentError, "need a server_uuid argument")}, nil
	}

	var action string
	var result string
	var errStr string
	var userID string
	var _time time.Time

	var db *sql.DB
	var rows *sql.Rows

	if _, err := os.Stat(dbFile(serverUUID)); os.IsNotExist(err) {
		err = errors.New("ShowServerActions(): Action log database file is not exist")
		goto ERROR
	}

	db, err = sql.Open("sqlite3", dbFile(serverUUID))
	if err != nil {
		goto ERROR
	}
	defer func() {
		_ = db.Close()
	}()

	rows, err = db.Query("SELECT * FROM server_actions ORDER BY time DESC")
	if err != nil {
		goto ERROR
	}

	for rows.Next() {
		err = rows.Scan(&action, &result, &errStr, &userID, &_time)
		if err != nil {
			goto ERROR
		}

		actions = append(actions, ServerAction{
			Action: action,
			Result: result,
			ErrStr: errStr,
			UserID: userID,
			Time:   _time,
		})
	}
	_ = rows.Close()

	serverActions.ServerActions = actions
	serverActions.Errors = hccerr.ReturnHccEmptyErrorPiccolo()

	return serverActions, nil

ERROR:
	serverActions.ServerActions = actions
	if err != nil {
		serverActions.Errors = hccerr.ReturnHccErrorPiccolo(hccerr.PiccoloInternalInitFail, err.Error())
	} else {
		serverActions.Errors = hccerr.ReturnHccEmptyErrorPiccolo()
	}

	return serverActions, nil
}
