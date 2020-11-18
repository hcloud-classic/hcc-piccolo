package mysql

import (
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/lib/logger"
	"testing"
)

func Test_DB_Prepare(t *testing.T) {
	err := logger.Init()
	if err != nil {
		errors.SetErrLogger(logger.Logger)
		errors.NewHccError(errors.PiccoloInternalInitFail, "logger.Init(): "+err.Error()).Fatal()
	}
	errors.SetErrLogger(logger.Logger)

	defer func() {
		logger.End()
	}()

	config.Init()

	err = Init()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		End()
	}()
}
