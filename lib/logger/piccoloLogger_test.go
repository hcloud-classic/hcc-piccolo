package logger

import (
<<<<<<< HEAD
	"hcc/piccolo/lib/syscheck"
=======
	"hcc/piccolo/lib/errors"
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	"testing"
)

func Test_CreateDirIfNotExist(t *testing.T) {
	err := CreateDirIfNotExist("/var/log/" + LogName)
	if err != nil {
		t.Fatal("Failed to create dir!")
	}
}

func Test_Logger_Prepare(t *testing.T) {
<<<<<<< HEAD
	err := syscheck.CheckRoot()
	if err != nil {
		t.Fatal("Failed to get root permission!")
	}

	if !Prepare() {
		t.Fatal("Failed to prepare logger!")
	}
	defer func() {
		_ = FpLog.Close()
=======
	err := Init()
	if err != nil {
		errors.SetErrLogger(Logger)
		errors.NewHccError(errors.PiccoloInternalInitFail, "logger.Init(): "+err.Error()).Fatal()
	}
	errors.SetErrLogger(Logger)

	defer func() {
		End()
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
	}()
}
