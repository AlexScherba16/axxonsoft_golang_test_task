package config

const (
	// server
	proxyServerPortEnvName = "PROXY_SERVER_PORT"
	defaultServerPort      = 8080

	maxHeaderBytesEnvName = "MAX_HEADER_BYTES"
	defaultMaxHeaderBytes = 1048576

	// timeout
	thirdPartyRequestTimeoutEnvName = "THIRD_PARTY_REQUEST_TIMEOUT_MS"
	defaultThirdPartyRequestTimeout = 1000

	serverReadTimeoutEnvName = "SERVER_READ_TIMEOUT_MS"
	defaultServerReadTimeout = 10000

	serverWriteTimeoutEnvName = "SERVER_WRITE_TIMEOUT_MS"
	defaultServerWriteTimeout = 10000
)
