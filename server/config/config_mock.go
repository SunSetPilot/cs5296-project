package config

func NewMockConfig() *Config {
	return &Config{
		DebugMode: true,
		AppName:   "cs5296-project-server-mock",
		ListenOn:  "127.0.0.1:8080",
		MySQL:     "mock",
		LogPath:   "./logs",
	}
}
