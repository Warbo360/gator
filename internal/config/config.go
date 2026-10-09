package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
    CurrentUsername string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {

	// Get the users home dir and error check

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	// Assuming no error, join the home dir with the filename for the full path to the config file and return that path

	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func Read() (Config, error) {

	// Get fullpath name for the config file and error check

	configPath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	// Read file at the built path and error check

	data, err := os.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	// Create empty config struct, unmarshal with a pointer to the new empty config struct, error check, and return if
	// there is no error.

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c Config) SetUser(username string) error {
	c.CurrentUsername = username
	return write(c)
}

func write(c Config) error {
	configPath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
