package config

type graphql struct {
	ProductionListenPort     int64 `goconf:"graphql:production_listen_port"`      // ProductionListenPort : Port number to receive GraphQL request through HTTP and websocket for production (No GraphiQL and Playground)
	DevInternalListenPort    int64 `goconf:"graphql:dev_internal_listen_port"`    // DevInternalListenPort : Port number to receive GraphQL request through HTTP and websocket for development
	DevInternalUsePlayground bool  `goconf:"graphql:dev_internal_use_playground"` // DevInternalUsePlayground : Use playground for GraphQL web UI for development
	SubscriptionInterval     int64 `goconf:"graphql:subscription_interval_ms"`    // SubscriptionInterval : Interval for GraphQL subscription
}

// GraphQL : graphql config structure
var GraphQL graphql
