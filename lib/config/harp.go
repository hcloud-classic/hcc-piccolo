package config

type harp struct {
	ServerAddress    string `goconf:"harp:harp_server_address"`     // ServerAddress : IP address of server which installed harp module
	ServerPort       int64  `goconf:"harp:harp_server_port"`        // ServerPort : Listening port number of harp module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"harp:harp_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to harp module
=======
	RequestTimeoutMs int64  `goconf:"harp:harp_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to harp module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// Harp : Harp config structure
var Harp harp
