package main

import "sync"

var instance *Config
var once sync.Once

type Config struct {
	DatabaseURL string
}

func GetConfigInstance() *Config {
	once.Do(func() {
		instance = &Config{DatabaseURL: "postgres://localhost"}
	})

	return instance
}
