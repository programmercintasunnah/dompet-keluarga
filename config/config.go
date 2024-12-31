package config

import (
	"os"
)

// GetEnv membaca variabel environment dengan default value jika nil.
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
