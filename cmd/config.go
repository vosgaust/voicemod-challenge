package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/storage/mysql"
)

type config struct {
	Host string `default:"localhost"`
	Port uint   `default:"8080"`

	DB mysql.Config
}

func getConfig() (*config, error) {
	cfg := config{}
	if err := envconfig.Process("voicemod", &cfg); err != nil {
		return nil, fmt.Errorf("failed to read the env: %v", err)
	}
	return &cfg, nil
}
