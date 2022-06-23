package config

import (
	"fmt"

	"github.com/gclamigueiro/sum-svc-gokit/internal/configurator"
)

func configEntries() []configurator.ConfigEntry {
	return []configurator.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Port used to listen for incoming connections",
			Shortcut:     "p",
			DefaultValue: "8081",
		},
	}
}

type APIConfig struct {
	Port string
}

func GetAPIConfig() *APIConfig {

	variables, err := configurator.LoadConfig(configEntries())

	if err != nil {
		fmt.Println(err)
	}

	return &APIConfig{
		Port: variables["port"].(string),
	}
}
