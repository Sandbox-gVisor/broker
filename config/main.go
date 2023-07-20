package config

func LoadConfig() Config {
	var (
		config Config
		err    error
	)
	// load from argparse
	config, err = parseCliArgs()
	if err == nil {
		return config
	}
	config, err = loadEnv()
	if err == nil {
		return config
	}
	config.Address = "localhost:9988"
	config.Type = "tcp"
	return config
}
