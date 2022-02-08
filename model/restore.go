package model

// Error -
type Error struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode string `json:"errcode"`
}

// Restore -
type Restore struct {
	RunStatus string `json:"runstatus"`
	RunUUID   string `json:"runuuid"`
	Errors    Error  `json:"errors"`
}

// RestoreData -
type RestoreData struct {
	Data struct {
		Subnet Restore `json:"restore"`
	} `json:"data"`
}
