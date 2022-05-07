package clihandler

type CLIHandlerParameters []*CLIHandlerParameter

type CLIHandlerParameter struct {
	Type         string
	Name         string
	Label        string
	DefaultValue string
	Examples     []string
	Validation   func(string) bool
}

func NewCLIHandlerParameter(type_ string, name string, label string,
	defaultValue string, examples []string, validation func(string) bool) *CLIHandlerParameter {

	if validation == nil {
		validation = func(string) bool { return true }
	}

	return &CLIHandlerParameter{
		Type:         type_,
		Name:         name,
		Label:        label,
		DefaultValue: defaultValue,
		Examples:     examples,
		Validation:   validation,
	}
}
