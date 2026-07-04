// Package config contains application-wide configuration constants.
package config

import "time"

// Authentication and session configuration.
const (
	SessionTimeout    = 30 * time.Minute
	MaxFailedAttempts = 5
	LockDuration      = 5 * time.Minute
)
