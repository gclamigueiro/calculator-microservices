package generator

import "fmt"

var registeredGenerators = make(map[string]Generator)

// interface that define the methods that a generator must implement to create a service template
type Generator interface {
	GetKey() string
	Process(outputPath string, parameters map[string]string) error
	ExcludedFiles(fullProjectPath string) map[string]int
}

/*type Parameter interface {
	GetName() string
	GetType() string
	GetRegex() string
}*/

// register a generator
func register(name string, generator Generator) {
	registeredGenerators[name] = generator
}

func GetRegisteredGeneratorsKeys() []string {
	keys := make([]string, 0, len(registeredGenerators))
	for k := range registeredGenerators {
		keys = append(keys, k)
	}
	return keys
}

func GetGenerator(name string) (Generator, error) {

	val, ok := registeredGenerators[name]

	if !ok {
		return nil, fmt.Errorf("generator %s not found", name)
	}

	return val, nil
}
