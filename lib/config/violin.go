package config

type violin struct {
	ServerAddress    string `goconf:"violin:violin_server_address"`     // ServerAddress : IP address of server which installed violin module
	ServerPort       int64  `goconf:"violin:violin_server_port"`        // ServerPort : Listening port number of violin module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"violin:violin_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to violin module
=======
	RequestTimeoutMs int64  `goconf:"violin:violin_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to violin module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// Violin : violin config structure
var Violin violin
