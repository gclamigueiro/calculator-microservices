package generator

import "fmt"

// registered Generators
var registeredGenerators = make(map[string]Generator)

// interface that define the methods that a generator must implement to create a service template
type Generator interface {
	GetKey() string
	Process(outputPath string, parameters map[string]string) error
	ExcludedFiles(fullProjectPath string) map[string]int
}

// register a generator
func register(name string, generator Generator) {
	registeredGenerators[name] = generator
}

// get the list of registered generators
func GetRegisteredGeneratorsKeys() []string {
	keys := make([]string, 0, len(registeredGenerators))
	for k := range registeredGenerators {
		keys = append(keys, k)
	}
	return keys
}

// get a generator
func GetGenerator(name string) (Generator, error) {

	val, ok := registeredGenerators[name]

	if !ok {
		return nil, fmt.Errorf("generator %s not found", name)
	}

	return val, nil
}
