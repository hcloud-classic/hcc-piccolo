package dao

import (
	dbsql "database/sql"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
	"strings"
)

// ReadGroup : Get infos of a group
func ReadGroup(groupID int) (*model.Group, error) {
	var group model.Group
	var groupName string

	sql := "select id, name from piccolo.group where id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row,
		&groupID,
		&groupName)
	if err != nil {
		errStr := "ReadGroup(): " + err.Error()
		logger.Logger.Println(errStr)

		return nil, err
	}

	group.ID = int64(groupID)
	group.Name = groupName

	return &group, nil
}

// ReadGroupList : Get list of groups
func ReadGroupList() (*pb.ResGetGroupList, uint64, string) {
	var groupList pb.ResGetGroupList
	var groups []pb.Group
	var pgroups []*pb.Group

	var id int64
	var name string

	sql := "select * from piccolo.group"

	var stmt *dbsql.Rows
	var err error
	stmt, err = mysql.Query(sql)

	if err != nil {
		errStr := "ReadGroupList(): " + err.Error()
		logger.Logger.Println(errStr)
		return nil, hcc_errors.PiccoloMySQLExecuteError, errStr
	}
	defer func() {
		_ = stmt.Close()
	}()

	for stmt.Next() {
		err := stmt.Scan(&id, &name)
		if err != nil {
			errStr := "ReadGroupList(): " + err.Error()
			logger.Logger.Println(errStr)
			if strings.Contains(err.Error(), "no rows in result set") {
				return nil, hcc_errors.PiccoloMySQLExecuteError, errStr
			}
			return nil, hcc_errors.PiccoloMySQLExecuteError, errStr
		}

		groups = append(groups, pb.Group{
			Id:   id,
			Name: name,
		})
	}

	for i := range groups {
		pgroups = append(pgroups, &groups[i])
	}

	groupList.Group = pgroups

	return &groupList, 0, ""
}
