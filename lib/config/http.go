package config

type http struct {
	Port          int64 `goconf:"http:port"`           // Port : Port number for receive graphql and websocket request via http server
	UsePlayground bool  `goconf:"http:use_playground"` // UsePlayground : Use playground for GraphQL web UI
}

// HTTP : http config structure
var HTTP http
