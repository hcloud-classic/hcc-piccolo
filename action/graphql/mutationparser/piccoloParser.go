package mutationparser

import (
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

	out, err := uuid.NewV4()
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloInternalUUIDGenerationError, err.Error())}, nil
	}
	UUID := out.String()

	user := model.User{
		UUID:     UUID,
		ID:       id,
		Password: password,
		Name:     name,
		Email:    email,
	}

	sql := "insert into user(uuid, id, password, name, email, created_at) values (?, ?, ?, ?, ?, now())"
	stmt, err := mysql.Db.Prepare(sql)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.UUID, user.ID, user.Password, user.Name, user.Email)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	return &user, nil
}
