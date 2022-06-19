package config

import "os"

// Config struct
type Config struct {
	ADDR string
}

// Get config
func GetConfig() *Config {
	return &Config{
		ADDR: getEnv("ADDR", ":8000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
