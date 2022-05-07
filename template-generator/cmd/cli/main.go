package main

import (
	"fmt"

	cliHandler "github.com/gclamigueiro/template-generator/internal/cli-handler"
	"github.com/gclamigueiro/template-generator/internal/generator"
)

func main() {

	parameters := cliHandler.CLIHandlerParameters{}
	servicesTypes := generator.GetRegisteredGeneratorsKeys()
	parameters = append(
		parameters,
		cliHandler.NewCLIHandlerParameter(
			"select",                               // type of input
			"type",                                 // parameter name
			"Select the template to generate",      // parameter label
			"",                                     // default value
			servicesTypes,                          // examples
			cliHandler.ValueInList(servicesTypes))) // validation

	parameters = append(
		parameters,
		cliHandler.NewCLIHandlerParameter(
			"text",                        // type of input
			"output_directory",            // parameter name
			"Input the output directory ", // parameter label
			"./output",                    // default value
			nil,                           // examples
			nil))                          // validation

	parameters = append(
		parameters,
		cliHandler.NewCLIHandlerParameter(
			"text",                           // type of input
			"APIName",                        // parameter name
			"Input the name of the service ", // parameter label
			"",                               // default value
			nil,                              // examples
			cliHandler.NotEmpty()))           // validation

	cliH := cliHandler.NewCliHandler(parameters)
	values := cliH.StartCLI()

	selectedGenerator, err := generator.GetGenerator(values["type"])

	if err != nil {
		fmt.Println(err)
	}

	err = selectedGenerator.Process(values["output_directory"], values)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("🎉 Project files generated 🎉")

}
