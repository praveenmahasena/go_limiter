// Package config ...
package config

import "time"

// Config ...
type Config struct {
	GlobalSettings   `json:"global_settings"`
	RateLimiter      `json:"rate_limiter"`
	Routing          `json:"routing"`
	ResponseHandling `json:"response_handling"`
}

// GlobalSettings ...
type GlobalSettings struct {
	Logging           bool          `json:"logging"`
	Metrics           bool          `json:"metrics"`
	CleanupIntervalms time.Duration `json:"cleanup_interval_ms"`
}

// RateLimiter ...
type RateLimiter struct {
	Rules []Rule `json:"rules"`
}

// Rule ...
type Rule struct {
	ID         string        `json:"id"`
	Path       string        `json:"path"`
	Limit      uint          `json:"limit"`
	Algorithm  string        `json:"algorithm"`
	Windowms   time.Duration `json:"window_ms"`
	HTTPMethod string        `json:"http_method"`
}

// Routing ...
type Routing struct {
	BackendURL    string `json:"backend_url"`
	GoLimiterPort string `json:"go_limiter_port"`
}

// ResponseHandling ...
type ResponseHandling struct {
	OnLimitExceeded `json:"on_limit_exceeded"`
}

// OnLimitExceeded ...
type OnLimitExceeded struct {
	HTTPStatus uint   `json:"http_status"`
	Message    string `json:"message"`
	RetryAfter bool   `json:"retry_after"`
}
