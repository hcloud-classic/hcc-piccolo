package config

type cello struct {
	ServerAddress    string `goconf:"cello:cello_server_address"`     // ServerAddress : IP address of server which installed cello module
	ServerPort       int64  `goconf:"cello:cello_server_port"`        // ServerPort : Listening port number of cello module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"cello:cello_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to cello module
=======
	RequestTimeoutMs int64  `goconf:"cello:cello_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to cello module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// Cello : cello config structure
var Cello cello
