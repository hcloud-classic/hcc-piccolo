package config

type tuba struct {
	ServerPort       int64 `goconf:"tuba:tuba_server_port"`        // ServerPort : Listening port number of tuba module
	RequestTimeoutMs int64 `goconf:"tuba:tuba_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to tuba module
}

// Tuba : Tuba config structure
var Tuba tuba
