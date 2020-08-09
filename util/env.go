package util

import "os"

// EVString return string type for environment variable by key. If it is not exist, return flashback
func EVString(key, flashback string) string {
	result := os.Getenv(key)

	if result == "" {
		return flashback
	}

	return key
}
