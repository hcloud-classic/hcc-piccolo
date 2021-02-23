package config

type graphql struct {
	Port                 int64 `goconf:"graphql:port"`                     // Port : Port number for receive GraphQL request through HTTP and websocket
	UsePlayground        bool  `goconf:"graphql:use_playground"`           // UsePlayground : Use playground for GraphQL web UI
	SubscriptionInterval int64 `goconf:"graphql:subscription_interval_ms"` // SubscriptionInterval : Interval for GraphQL subscription
}

// GraphQL : graphql config structure
var GraphQL graphql
