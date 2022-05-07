package generator

import (
	"fmt"
	"path"

	templateHandler "github.com/gclamigueiro/template-generator/internal/template-handler"
)

func init() {
	register("go-kit", &GoKitGenerator{})
}

type GoKitGenerator struct {
}

func (g *GoKitGenerator) GetKey() string {
	return "go-kit"
}

func (g *GoKitGenerator) Process(outputPath string, parameters map[string]string) error {
	fmt.Println("Processing Go Kit Generator")

	templateProcessor := templateHandler.NewFileProcessor()

	fullTemplatePath := path.Join("templates", g.GetKey())
	fullProjectPath := path.Join(outputPath, parameters["APIName"])

	err := templateProcessor.CreateInitialDirectory(fullTemplatePath, fullProjectPath)

	if err != nil {
		return err
	}

	if err := templateProcessor.ProcessMultipleTemplates(fullProjectPath, parameters, g.ExcludedFiles(fullProjectPath)); err != nil {
		return err
	}

	return nil
}

func (g *GoKitGenerator) ExcludedFiles(fullProjectPath string) map[string]int {
	return map[string]int{
		fullProjectPath + "/iaas/helm/templates": 1,
	}
}
