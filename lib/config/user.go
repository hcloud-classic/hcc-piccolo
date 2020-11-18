package config

type user struct {
	TokenExpirationTimeMinutes int64 `goconf:"user:token_expiration_time_minutes"` // TokenExpirationTimeMinutes : Expiration time of the token (minutes)
}

// User : user config structure
var User user
