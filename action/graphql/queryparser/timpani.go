package queryparser

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/model"
	"os/exec"
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
