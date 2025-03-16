package main

import (
	"os"
)

// GetEnv gets the environment variable from the .env file, accepting a default if not found
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
