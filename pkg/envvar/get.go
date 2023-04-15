package envvar

import (
	"os"
)

// GetEnv is a os.Getenv() wrapper adding a default value
func GetEnv(envVar string, defaultValue string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return defaultValue
	}
	return val
}
