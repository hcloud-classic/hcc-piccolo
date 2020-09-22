package queryparser

import (
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcflute"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/usertool"
	"hcc/piccolo/model"
)

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

	token, err := usertool.GenerateToken(id, password)
	if err != nil {
		return model.Token{Token: "", Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLTokenGenerationError, err.Error())}, nil
	}

	return model.Token{Token: token, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
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
