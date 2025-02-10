package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db     *DB
	Server *Server
}

type Server struct {
	Port int `yaml:"port"`
}

type DB struct {
	Url string `yaml:"url"`
}

var (
	once           sync.Once
	configInstance Config
)

func GetConfig() Config {
	once.Do(func() {
		configPath := "config.yml"

		f, err := os.Open(configPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&configInstance)

		if err != nil {
			panic(err)
		}
	})

	return configInstance
}
