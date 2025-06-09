// Package config ...
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load loades up the config
func Load() (*Config, error) {
	path, pathErr := getPath()
	if pathErr != nil {
		return nil, pathErr
	}
	config, configErr := conf(path)
	if configErr != nil {
		return nil, fmt.Errorf("error during reading out config file")
	}
	return config, nil
}

func getPath() (string, error) {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		return "", fmt.Errorf("error during getting working dir with value %v", wdErr)
	}
	return wd + "/config.json", nil
}

func conf(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error during reading out config file %v", err)
	}
	conf := &Config{}
	if err := json.Unmarshal(b, conf); err != nil {
		return nil, fmt.Errorf("error during writing to config with value %v", err)
	}
	ServerAddr = conf.BackendURL
	return conf, nil
}
