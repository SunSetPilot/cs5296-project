package config

type Config struct {
	DebugMode bool
	AppName   string
	ListenOn  string
	MySQL     string
	LogPath   string
}

func MustLoad(configFile string, config *Config) {
	return
}
