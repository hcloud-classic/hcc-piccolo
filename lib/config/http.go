package config

type http struct {
<<<<<<< HEAD
	Port             int64 `goconf:"http:port"`               // Port : Port number for listening graphql request via http server
	RequestTimeoutMs int64 `goconf:"http:request_timeout_ms"` // RequestTimeoutMs : Timeout for HTTP request
=======
	Port          int64 `goconf:"http:port"`           // Port : Port number for receive graphql and websocket request via http server
	UsePlayground bool  `goconf:"http:use_playground"` // UsePlayground : Use playground for GraphQL web UI
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// HTTP : http config structure
var HTTP http
