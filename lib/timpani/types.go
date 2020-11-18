package timpani

type normalRebootNotificationRequestType struct {
	NodeUUID string `json:"node_uuid"`
	BootKind string `json:"boot_kind"`
}

type normalRebootNotificationResponseType struct {
	Result        string `json:"result"`
	ResultMessage string `json:"resultMessage"`
	ResultData    struct {
		BootKind string `json:"boot_kind"`
	} `json:"resultData"`
}
