package config

func getDefaultConfig() Config {
	return Config{
		Address:       "localhost:12001",
		Type:          "tcp",
		WebsoketPort:  "8080",
		QueueName:     "main",
		RabbitAddress: "amqp://guest:guest@localhost:5672/",
	}
}

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
	// return default config
	return getDefaultConfig()
}
