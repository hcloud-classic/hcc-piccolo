package config

type violinNoVnc struct {
	ServerAddress    string `goconf:"violinnovnc:violinnovnc_server_address"`     // ServerAddress : IP address of server which installed violin-novnc module
	ServerPort       int64  `goconf:"violinnovnc:violinnovnc_server_port"`        // ServerPort : Listening port number of violin-novnc module
	RequestTimeoutMs int64  `goconf:"violinnovnc:violinnovnc_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to violin-novnc module
}

// ViolinNoVnc : violinNoVnc config structure
var ViolinNoVnc violinNoVnc
