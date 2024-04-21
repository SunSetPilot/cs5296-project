package config

func NewMockConfig() *Config {
	return &Config{
		DebugMode:         true,
		AppName:           "cs5296-project-client-mock",
		LogPath:           "./logs",
		HeartbeatInterval: 5,
		FetchTaskInterval: 5,
		ExecPoolSize:      5,
	}
}
