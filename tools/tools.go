package tools

import (
	"time"
)

// GetTime is a function that returns the current time
func GetTime() string {
	return time.Now().Format(time.RFC3339)
}
