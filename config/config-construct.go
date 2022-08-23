package config

import (
	"os"
	"strconv"
)

// NewApplicationSettings - applications context settings constructor
func NewApplicationSettings() *ApplicationSettings {
	return &ApplicationSettings{
		Server:  newServerConfig(),
		Timeout: newTimeOutConfig(),
	}
}

// newServerConfig - server context settings private constructor
func newServerConfig() ServerConfig {
	return ServerConfig{
		Port:           getEnvAsInt(proxyServerPortEnvName, defaultServerPort),
		MaxHeaderBytes: getEnvAsInt(maxHeaderBytesEnvName, defaultMaxHeaderBytes),
	}
}

// newTimeOutConfig - timeout context settings private constructor
func newTimeOutConfig() TimeOutConfig {
	return TimeOutConfig{
		ThirdPartyRequestTimeoutMs: getEnvAsInt(thirdPartyRequestTimeoutEnvName, defaultThirdPartyRequestTimeout),
		ServerReadTimeoutMs:        getEnvAsInt(serverReadTimeoutEnvName, defaultServerReadTimeout),
		ServerWriteTimeoutMs:       getEnvAsInt(serverWriteTimeoutEnvName, defaultServerWriteTimeout),
	}
}

// getEnv - read environment value by name and convert to string
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvAsInt - read and convert environment value to int
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
