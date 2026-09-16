package config

import "time"

type Config struct {
	Host string
	Port int

	ShutdownTime time.Duration

	GameTickInterval time.Duration
}

func Default() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,

		ShutdownTime: 10 * time.Second,

		GameTickInterval: time.Second / 60,
	}
}
