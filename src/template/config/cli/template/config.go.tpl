package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port int    `json:"port"`
	Env  string `json:"env"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
