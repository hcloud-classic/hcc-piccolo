package pid

import (
	"hcc/piccolo/lib/fileutil"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

var piccoloPIDFileLocation = "/var/run"
var piccoloPIDFile = "/var/run/piccolo.pid"

// IsPiccoloRunning : Check if piccolo is running
func IsPiccoloRunning() (running bool, pid int, err error) {
	if _, err := os.Stat(piccoloPIDFile); os.IsNotExist(err) {
		return false, 0, nil
	}

	pidStr, err := ioutil.ReadFile(piccoloPIDFile)
	if err != nil {
		return false, 0, err
	}

	piccoloPID, _ := strconv.Atoi(string(pidStr))

	proc, err := os.FindProcess(piccoloPID)
	if err != nil {
		return false, 0, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, piccoloPID, nil
	}

	return false, 0, nil
}

// WritePiccoloPID : Write piccolo PID to "/var/run/piccolo.pid"
func WritePiccoloPID() error {
	pid := os.Getpid()

	err := fileutil.CreateDirIfNotExist(piccoloPIDFileLocation)
	if err != nil {
		return err
	}

	err = fileutil.WriteFile(piccoloPIDFile, strconv.Itoa(pid))
	if err != nil {
		return err
	}

	return nil
}

// DeletePiccoloPID : Delete the piccolo PID file
func DeletePiccoloPID() error {
	err := fileutil.DeleteFile(piccoloPIDFile)
	if err != nil {
		return err
	}

	return nil
}
