package config

import "os"

type Config struct {
	Port             string
	LogLevel         string
	TenantID         string
	ClientID         string
	ExpectedAudience string
}

func Load() Config {
	return Config{
		Port:             getEnv("PORT", "8080"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
		TenantID:         getEnv("AZURE_TENANT_ID", ""),
		ClientID:         getEnv("AZURE_CLIENT_ID", ""),
		ExpectedAudience: getEnv("AZURE_EXPECTED_AUDIENCE", "api://entrax-api"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
