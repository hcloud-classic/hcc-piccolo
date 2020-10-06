package mutationparser

import (
	"crypto/sha256"
	"encoding/hex"
	uuid "github.com/nu7hatch/gouuid"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
)

// SignUp : Do user sign up process
func SignUp(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)
	password, passwordOk := args["password"].(string)
	name, nameOk := args["name"].(string)
	email, emailOk := args["email"].(string)

	if !idOk || !passwordOk || !nameOk || !emailOk {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need id and password, name, email arguments")}, nil
	}

	sql := "select id from user where id = ?"
	err := mysql.Db.QueryRow(sql, id).Scan(&id)
	if err == nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLUserExist, "Provided ID is in use")}, nil
	}

	out, err := uuid.NewV4()
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloInternalUUIDGenerationError, err.Error())}, nil
	}
	UUID := out.String()

	hash := sha256.New()
	hash.Write([]byte(password))
	hashPassword := hex.EncodeToString(hash.Sum(nil))

	user := model.User{
		UUID:     UUID,
		ID:       id,
		Name:     name,
		Email:    email,
	}

	sql = "insert into user(uuid, id, password, name, email, login_at, created_at) values (?, ?, ?, ?, ?, now(), now())"
	stmt, err := mysql.Db.Prepare(sql)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.UUID, user.ID, hashPassword, user.Name, user.Email)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user.Errors = errors.ReturnHccEmptyErrorPiccolo()

	return &user, nil
}
