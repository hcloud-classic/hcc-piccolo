package model

import "hcc/piccolo/action/grpc/errconv"

type ErrorField struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type TimpaniService struct {
	Target string                    `json:"target"`
	Result string                    `json:"result"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}

type MasterSync struct {
	Data struct {
		Isvaild bool       `json:"isvaild"`
		Errors  ErrorField `json:"errors"`
	} `json:"data"`
}

// Backup : CmdResponse
type Backup struct {
	Data struct {
		Runstatus string     `json:"runstatus"`
		RunUUID   string     `json:"runuuid"`
		Errors    ErrorField `json:"errors"`
	} `json:"data"`
}

// BackupScheduler : CmdResponse
type BackupScheduler struct {
	Data struct {
		BackupScheduler struct {
			Status string     `json:"status"`
			Errors ErrorField `json:"errors"`
		} `json:"backupScheduler"`
	} `json:"data"`
}
