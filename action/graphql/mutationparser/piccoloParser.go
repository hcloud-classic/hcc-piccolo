package mutationparser

import (
	uuid "github.com/nu7hatch/gouuid"
	"hcc/piccolo/action/graphql/queryparser"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"strings"
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

	if strings.ToLower(id) == "admin" || strings.ToLower(id) == "administrator" {
		logger.Logger.Println("SignUp(): Someone tried to sign up with one of administrative ID.")
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLUserExist, "Hey, hey you! Yeah, you.")}, nil
	}

	sql := "select id from user where id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &id)
	if err == nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLUserExist, "Provided ID is in use")}, nil
	}

	out, err := uuid.NewV4()
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloInternalUUIDGenerationError, err.Error())}, nil
	}
	UUID := out.String()

	user := model.User{
		UUID:  UUID,
		ID:    id,
		Name:  name,
		Email: email,
	}

	sql = "insert into user(uuid, id, password, name, email, login_at, created_at) values (?, ?, ?, ?, ?, now(), now())"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.UUID, user.ID, password, user.Name, user.Email)
	if err != nil {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user.Errors = errors.ReturnHccEmptyErrorPiccolo()

	return &user, nil
}

// Unregister : Do user unregister process
func Unregister(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)

	if !idOk {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "need a id argument")}, nil
	}

	if strings.ToLower(id) == "admin" || strings.ToLower(id) == "administrator" {
		logger.Logger.Println("Unregister(): Someone tried to unregister one of administrative ID.")
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLUserExist, "You can't delete administrative IDs")}, nil
	}

	user, _ := queryparser.User(args)
	if len(user.(model.User).Errors) != 0 && user.(model.User).Errors[0].ErrCode != 0 {
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, "user not found")}, nil
	}

	sql := "delete from user where id = ?"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "Unregister(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		errStr := "Unregister(): " + err2.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	return model.User{ID: id, Errors: errors.ReturnHccEmptyErrorPiccolo()}, nil
}
