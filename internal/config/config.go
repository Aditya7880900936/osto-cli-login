package config

import "time"

const (
	SessionTimeout    = 30 * time.Minute
	MaxFailedAttempts = 5
	LockDuration      = 5 * time.Minute
)