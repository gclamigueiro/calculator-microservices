package config

import (
	"fmt"

	"multiple-operation-svc-go-kit/internal/configurator"
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
			VariableName: "uri_sum_service",
			Description:  "Url of the Sum Service",
			DefaultValue: "http://localhost:8081/",
		},
		{
			VariableName: "uri_subtract_service",
			Description:  "Url of the Subtract Service",
			DefaultValue: "http://localhost:8082/",
		},
	}
}

type APIConfig struct {
	Port               string
	UriSumService      string
	UriSubtractService string
}

func GetAPIConfig() *APIConfig {

	variables, err := configurator.LoadConfig(configEntries())

	if err != nil {
		fmt.Println(err)
	}

	return &APIConfig{
		Port:               variables["port"].(string),
		UriSumService:      variables["uri_sum_service"].(string),
		UriSubtractService: variables["uri_subtract_service"].(string),
	}
}
