package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	json_data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(json_data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
