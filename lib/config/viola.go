package config

type viola struct {
	ServerAddress    string `goconf:"viola:viola_server_address"`     // ServerAddress : IP address of server which installed viola module
	ServerPort       int64  `goconf:"viola:viola_server_port"`        // ServerPort : Listening port number of viola module
	RequestTimeoutMs int64  `goconf:"viola:viola_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to viola module
}

// Viola : viola config structure
var Viola viola
