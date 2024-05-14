package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ServerMap struct{ Prefix, Server string }

type Config struct{ Mappings []ServerMap }

func LoadConfig() (*Config, error) {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	absConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		log.Fatalf("failed to get absolute path for config file: %v", err)
	}

	fmt.Println("Using config file:", absConfigPath)

	if _, err := os.Stat(absConfigPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", absConfigPath)
	}

	data, err := os.ReadFile(absConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	return &config, nil
}
