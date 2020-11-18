package config

type violinNoVnc struct {
	ServerAddress    string `goconf:"violinnovnc:violinnovnc_server_address"`     // ServerAddress : IP address of server which installed violin-novnc module
	ServerPort       int64  `goconf:"violinnovnc:violinnovnc_server_port"`        // ServerPort : Listening port number of violin-novnc module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"violinnovnc:violinnovnc_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to violin-novnc module
=======
	RequestTimeoutMs int64  `goconf:"violinnovnc:violinnovnc_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to violin-novnc module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// ViolinNoVnc : violinNoVnc config structure
var ViolinNoVnc violinNoVnc
