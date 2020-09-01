package mutationparser

import (
	uuid "github.com/nu7hatch/gouuid"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
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
		return nil, errors.NewHccError(errors.PiccoloGraphQLArgumentError, "need id and password, name, email arguments").New()
	}

	out, err := uuid.NewV4()
	if err != nil {
		errors.NewHccError(errors.PiccoloInternalUUIDGenerationError, err.Error()).Println()
		return nil, err
	}
	UUID := out.String()

	user := model.User{
		UUID:     UUID,
		Id:       id,
		Password: password,
		Name:     name,
		Email:    email,
	}

	sql := "insert into user(uuid, id, password, name, email, created_at) values (?, ?, ?, ?, ?, now())"
	stmt, err := mysql.Db.Prepare(sql)
	if err != nil {
		errors.NewHccError(errors.PiccoloMySQLPrepareError, err.Error()).Println()
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()
	result, err := stmt.Exec(user.UUID, user.Id, user.Password, user.Name, user.Email)
	if err != nil {
		errors.NewHccError(errors.PiccoloMySQLExecuteError, err.Error()).Println()
		return nil, err
	}
	logger.Logger.Println(result.LastInsertId())

	return &user, nil
}
