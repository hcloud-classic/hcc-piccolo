package queryparser

import (
	"fmt"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/http/graphqlreq"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/model"
	"os/exec"
	"strconv"
	"strings"

	"innogrid.com/hcloud-classic/hcc_errors"
)

func TimapniServiceController(args map[string]interface{}) (interface{}, error) {
	// var (
	// 	serviceName string
	// 	action      string
	// )
	var (
		Service model.Service
		result  bool
	)
	target, targetOK := args["target"].(string)
	action, actionOK := args["action"].(string)
	if !targetOK {
		goto ERROR
	}
	Service.Target = target
	if !actionOK {
		goto ERROR
	}
	Service.Action = action
	result, _ = TimpaniServiceControlWithSSH(target, action)

	if !result {
		goto ERROR
	}
	Service.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	Service.Status = "true"
	return Service, nil
ERROR:
	Service.Status = "Fail"
	return model.Service{Target: target, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Failed")}, nil
}

func TimpaniServiceControlWithSSH(target string, action string) (bool, string) {
	var ip string
	if strings.ToLower(target) == "master" {
		ip = config.Mysql.Address
	} else if strings.ToLower(target) == "storage" {
		ip = config.Cello.ServerAddress
	} else {
		return false, "Please Choose master/storage"
	}
	if !(len(action) > 0) {
		return false, "Please Correct action start/restart/stop,status"
	}
	serviceRestart := "ssh -o StrictHostKeyChecking=no -T root@" + ip + " -p 22 \"service timpani " + strings.ToLower(action) + "\""
	cmd := exec.Command("/bin/sh", "-c", serviceRestart)
	result, err := cmd.CombinedOutput()
	if err != nil {
		logger.Logger.Println(ip + " Timpani Service Can't start")
		return false, err.Error()
	}
	// }else {
	// 	logger.Logger.Println(ip + " Telegraf Service Already started")
	// }
	return true, strings.TrimSuffix(string(result), "\n")
}

func TimpaniMasterSync(args map[string]interface{}) (interface{}, error) {
	var masterSync model.MasterSync
	query := "query {\n" +
		"	mastersync {\n" +
		"		token\n" +
		"		username\n" +
		"		newpw\n" +
		"	}\n" +
		"}"
	result, err := graphqlreq.DoHTTPRequest("timpani", true, masterSync, "mastersync", query)
	if err != nil {
		goto ERROR
	}
	return result, nil

ERROR:
	masterSync.Data.Errors.Errmsg = "TimpaniMasterSync Failed " + err.Error()
	// return model.Service{Target: target, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Failed")}, nil
	return masterSync, err
}

func TimpaniBackup(args map[string]interface{}) (interface{}, error) {
	var volumeBackup model.Backup
	query := "query {\n" +
		"	backup {\n" +
		"		uuid\n" +
		"		usetype\n" +
		"		nodetype\n" +
		"		name\n" +
		"	}\n" +
		"}"
	result, err := graphqlreq.DoHTTPRequest("timpani", true, volumeBackup, "backup", query)
	if err != nil {
		goto ERROR
	}
	return result, nil

ERROR:
	volumeBackup.Data.Errors.Errmsg = "TimpaniBackup Failed " + err.Error()
	// return model.Service{Target: target, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Failed")}, nil
	return volumeBackup, err
}

func TimpaniBackupScheduler(args map[string]interface{}) (interface{}, error) {
	var volumeBackup model.Backup
	query := "query {\n" +
		"	backup {\n" +
		"		uuid\n" +
		"		usetype\n" +
		"		nodetype\n" +
		"		name\n" +
		"	}\n" +
		"}"
	result, err := graphqlreq.DoHTTPRequest("timpani", true, volumeBackup, "backup", query)
	if err != nil {
		goto ERROR
	}
	return result, nil

ERROR:
	volumeBackup.Data.Errors.Errmsg = "TimpaniBackup Failed " + err.Error()
	// return model.Service{Target: target, Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Failed")}, nil
	return volumeBackup, err
}

func Restore(args map[string]interface{}) (interface{}, error) {
	snapname, _ := args["snapname"].(string)
	usetype, _ := args["usetype"].(string)
	nodetype, _ := args["nodetype"].(string)
	isboot, _ := args["isboot"].(bool)
	token, _ := args["token"].(string)

	var restoreData model.RestoreData
	query := "query { restore (snapname: \"" + snapname + "\", usetype: \"" + usetype + "\", nodetype: \"" + nodetype + "\", " +
		"isboot: " + strconv.FormatBool(isboot) + ", token: \"" + token + "\")" +
		"{ runstatus runuuid errors { errmsg errcode } } } }"
	fmt.Println(query)

	return graphqlreq.DoHTTPRequest("timpani", true, restoreData, "restore", query)
}
