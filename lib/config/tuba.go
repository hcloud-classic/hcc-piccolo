package config

type tuba struct {
	RequestTimeoutMs int64 `goconf:"tuba:tuba_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to tuba module
}

// Tuba : Tuba config structure
var Tuba tuba
