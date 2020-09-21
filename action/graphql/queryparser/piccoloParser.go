package queryparser

import (
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/userTool"
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

	token, err := userTool.GenerateToken(id, password)
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

	err := userTool.ValidateToken(args)
	if err != nil {
		return model.IsValid{IsValid: false}, nil
	}

	return model.IsValid{IsValid: true}, nil
}
