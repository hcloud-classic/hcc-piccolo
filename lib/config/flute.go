package config

type flute struct {
	ServerAddress    string `goconf:"flute:flute_server_address"`     // ServerAddress : IP address of server which installed flute module
	ServerPort       int64  `goconf:"flute:flute_server_port"`        // ServerPort : Listening port number of flute module
<<<<<<< HEAD
	RequestTimeoutMs int64  `goconf:"flute:flute_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for GraphQL request to flute module
=======
	RequestTimeoutMs int64  `goconf:"flute:flute_request_timeout_ms"` // RequestTimeoutMs : HTTP timeout for gRPC request to flute module
>>>>>>> eebb5a0417798d0031b913a3fa3db7ac18f22d33
}

// Flute : flute config structure
var Flute flute
