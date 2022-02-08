package model

type ErrorField struct {
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type TimpaniService struct {
	Target string     `json:"target"`
	Result string     `json:"result"`
	Errors ErrorField `json:"errors"`
}

// Backup -
type Backup struct {
	RunStatus string     `json:"runstatus"`
	RunUUID   string     `json:"runuuid"`
	Errors    ErrorField `json:"errors"`
}

// Restore -
type Restore struct {
	RunStatus string     `json:"runstatus"`
	RunUUID   string     `json:"runuuid"`
	Errors    ErrorField `json:"errors"`
}

type MasterSync struct {
	Data struct {
		Isvaild bool       `json:"isvaild"`
		Errors  ErrorField `json:"errors"`
	} `json:"data"`
}

// BackupData : CmdResponse
type BackupData struct {
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

// RestoreData -
type RestoreData struct {
	Data struct {
		Restore Restore `json:"restore"`
	} `json:"data"`
}
