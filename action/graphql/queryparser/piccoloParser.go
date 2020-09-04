package queryparser

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/userTool"
	"hcc/piccolo/model"
)

var loginMismatchError = errors.NewHccError(errors.PiccoloGraphQLLoginFailed, "user not found or password mismatch")

// Login : Do user login process
func Login(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)
	password, passwordOk := args["password"].(string)

	fmt.Println(*errors.NewHccErrorStack(errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need id and password arguments")).ConvertReportForm())

	if !idOk || !passwordOk {
		return model.Token{Token: "", Errors: *errors.NewHccErrorStack(errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need id and password arguments")).ConvertReportForm()}, nil
	}

	var dbPassword string

	sql := "select password from user where id = ?"
	err := mysql.Db.QueryRow(sql, id).Scan(&dbPassword)
	if err != nil {
		logger.Logger.Println(err)

		return model.Token{Token: "", Errors: *errors.NewHccErrorStack(loginMismatchError).ConvertReportForm()}, nil
	}

	// Given password is hashed password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
	if err != nil {
		return model.Token{Token: "", Errors: *errors.NewHccErrorStack(loginMismatchError).ConvertReportForm()}, nil
	}

	logger.Logger.Println("User logged in: " + id)

	token, err := userTool.GenerateToken(id, password)
	if err != nil {
		return model.Token{Token: "", Errors: *errors.NewHccErrorStack(errors.NewHccError(errors.PiccoloGraphQLTokenGenerationError, err.Error())).ConvertReportForm()}, nil
	}

	return model.Token{Token: token, Errors: *errors.NewHccErrorStack()}, nil
}

// CheckToken : Do token validation check process
func CheckToken(args map[string]interface{}) (interface{}, error) {
	token, tokenOk := args["token"].(string)

	if !tokenOk {
		return model.IsValid{IsValid: false, Errors: *errors.NewHccErrorStack(errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need a token argument")).ConvertReportForm()}, nil
	}

	var tokenArg = make(map[string]interface{})
	tokenArg["token"] = token
	err := userTool.ValidateToken(tokenArg)
	if err != nil {
		return model.IsValid{IsValid: false}, nil
	}

	return model.IsValid{IsValid: true}, nil
}
