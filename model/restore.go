package model

// Error -
type Error struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode string `json:"errcode"`
}

// RestoreInfo -
type RestoreInfo struct {
	RunStatus string `json:"runstatus"`
	RunUUID   string `json:"runuuid"`
	Errors    Error  `json:"errors"`
}

// Restore -
type Restore struct {
	RestoreInfo RestoreInfo `json:"restore"`
}

// RestoreData -
type RestoreData struct {
	Data struct {
		Subnet Restore `json:"restore"`
	} `json:"data"`
}
