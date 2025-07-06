package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	json_data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, json_data, 0600)
	if err != nil {
		return err
	}

	return nil
}
