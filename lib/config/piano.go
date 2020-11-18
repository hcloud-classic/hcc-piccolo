package config

type piano struct {
	ServerAddress    string `goconf:"piano:pianon_server_address"`    // ServerAddress : IP address of server which installed piano module
	ServerPort       int64  `goconf:"piano:piano_server_port"`        // ServerPort : Listening port number of piano module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"piano:piano_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to piano module
=======
	RequestTimeoutMs int64  `goconf:"piano:piano_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to piano module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// Piano : piano config structure
var Piano piano
