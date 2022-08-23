package config

// ServerConfig - server context settings
type ServerConfig struct {
	Port           int
	MaxHeaderBytes int
	// add future server settings here
}

// TimeOutConfig - timeouts context settings
type TimeOutConfig struct {
	ThirdPartyRequestTimeoutMs int
	ServerReadTimeoutMs        int
	ServerWriteTimeoutMs       int
	// add future timeout settings here
}

// ApplicationSettings - applications context settings
type ApplicationSettings struct {
	Server  ServerConfig
	Timeout TimeOutConfig
}
