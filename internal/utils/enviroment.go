package utils

import "os"

// GetEnv allows retrieving environment variables with fallback to default values
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
