package queryparser

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/lib/userTool"
	"hcc/piccolo/model"
)

var loginMismatchError = errors.New("user not found or password mismatch")

// Login : Do user login process
func Login(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)
	password, passwordOk := args["password"].(string)

	if !idOk || !passwordOk {
		return nil, errors.New("need id and password arguments")
	}

	var dbPassword string

	sql := "select password from user where id = ?"
	err := mysql.Db.QueryRow(sql, id).Scan(&dbPassword)
	if err != nil {
		logger.Logger.Println(err)
		return nil, loginMismatchError
	}

	// Given password is hashed password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
	if err != nil {
		return nil, loginMismatchError
	}

	logger.Logger.Println("User logged in: " + id)

	token, err := userTool.GenerateToken(id, password)
	if err != nil {
		return nil, err
	}

	return model.Token{Token: token}, nil
}
