package config

import (
	"fmt"

	"{{.APINamespace}}{{.APIName}}/internal/configurator"
)

func configEntries() []configurator.ConfigEntry {
	return []configurator.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Port used to listen for incoming connections",
			Shortcut:     "p",
			DefaultValue: "8080",
		},
		{
			VariableName: "uri_service",
			Description:  "Url of callable service",
			DefaultValue: "http://localhost:8081/",
		},
	}
}

type APIConfig struct {
	Port          string
	UriService string
}

func GetAPIConfig() *APIConfig {

	variables, err := configurator.LoadConfig(configEntries())

	if err != nil {
		fmt.Println(err)
	}

	return &APIConfig{
		Port:          variables["port"].(string),
		UriService: variables["uri_service"].(string),
	}
}
