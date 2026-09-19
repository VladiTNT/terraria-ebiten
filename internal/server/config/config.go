package config

import "time"

type Config struct {
	Host string
	Port int

	ShutdownTime time.Duration

	GameSettings *Settings
}

func Default() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,

		ShutdownTime: 10 * time.Second,

		GameSettings: &Settings{},
	}
}
