package clihandler

import (
	"fmt"
)

// before execute the Generator, the CLI is executed to get the parameters from the user
type CLIHandler interface {
	StartCLI() map[string]string
}

type cliHandler struct {
	parameters CLIHandlerParameters
}

func NewCliHandler(parameters CLIHandlerParameters) CLIHandler {
	return &cliHandler{
		parameters: parameters,
	}
}

func (cliHandler *cliHandler) StartCLI() map[string]string {

	results := make(map[string]string)

	for _, parameter := range cliHandler.parameters {

		var result = ""
		print(parameter, &result)

		for !parameter.Validation(result) {
			print(parameter, &result)
		}

		if result == "" && parameter.DefaultValue != "" {
			result = parameter.DefaultValue
		}

		results[parameter.Name] = result
	}

	return results
}

func print(parameter *CLIHandlerParameter, result *string) {

	text := ""

	if parameter.Examples != nil && len(parameter.Examples) > 0 {

		if parameter.Type == "text" {
			text = fmt.Sprintf("%s(examples: %s)", parameter.Label, parameter.Examples[0])
		} else {
			text = fmt.Sprintf("%s(accepted values: %s)", parameter.Label, parameter.Examples[0])
		}

	} else {
		text = parameter.Label
	}

	if parameter.DefaultValue != "" {
		text = fmt.Sprintf("%s[%s]:", text, parameter.DefaultValue)
	} else {
		text = text + ":"
	}

	fmt.Print(text)

	fmt.Scanln(result)
}
