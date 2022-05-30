package configurator

import (
	"errors"

	"github.com/spf13/viper"
)

type ConfigEntry struct {
	VariableName string
	Description  string
	Shortcut     string
	DefaultValue interface{}
}

func LoadConfig(entries []ConfigEntry) (map[string]interface{}, error) {

	if len(entries) == 0 {
		return nil, errors.New("no entries to configure")
	}

	//Configuration for enviromental variables on the system
	viper.AutomaticEnv()

	values := make(map[string]interface{})

	for _, entry := range entries {

		name := entry.VariableName

		val := viper.Get(name)
		if val == nil {
			val = entry.DefaultValue
		}

		values[name] = val
	}
	return values, nil
}
