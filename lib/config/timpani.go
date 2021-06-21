package config

type timpani struct {
	ServerAddress    string `goconf:"timpani:timpani_server_address"`     // ServerAddress : IP address of server which installed timpani module
	ServerPort       int64  `goconf:"timpani:timpani_server_port"`        // ServerPort : Listening port number of timpani module
	RequestTimeoutMs int64  `goconf:"timpani:timpani_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to timpani module
	RequestRetry     int64  `goconf:"timpani:timpani_request_retry"`      // RequestRetry : HTTP retry counts for GraphQL request to timpani module
}

// Timpani : timpani config structure
var Timpani timpani
