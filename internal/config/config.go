package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database    Database    `json:"database"`
	Application Application `json:"application"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Application struct {
	Port string `json:"port"`
}

// GetConfig get the config given a path
func GetConfig(configPath string) (Config, error) {
	// attempt to load config from config file path
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
