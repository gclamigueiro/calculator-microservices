package template_handler

import (
	"bytes"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	ErrParseFile                = "Archivo no existe o path invalido"
	ErrNilContext               = "Contexto no debe ser nil"
	ErrTemplateExecute          = "Error ejecutando template"
	ErrCreatingFileFromTemplate = "Error creando archivo desde template"
	ErrWritingFileFromTemplate  = "Error escribiendo archivo desde template"
)

// FileProcessor interfaz de procesador de archivos de plantillas
type FileProcessor interface {
	CreateInitialDirectory(src, dest string) error
	ProcessMultipleTemplates(basePath string, context interface{}, excludedFiles map[string]int) error
	ApplyTemplateToFile(workPath string, templateContext interface{}) error
	ProcessFile(path string, context interface{}) (string, error)
}

type fileProcessor struct{}

// NewFileProcessor crea un nuevo template processor
func NewFileProcessor() FileProcessor {
	return &fileProcessor{}
}

func (f *fileProcessor) CreateInitialDirectory(src, dest string) error {

	// First, remove the destination directory in case exists
	os.RemoveAll(dest)

	if err := os.MkdirAll(dest, os.FileMode(0764)); err != nil {
		return errors.New("Error creando directorio de proyecto" + " // " + err.Error())
	}

	CopyDir(src, dest)

	return nil
}

func (f *fileProcessor) ProcessMultipleTemplates(basePath string, context interface{}, excludedFiles map[string]int) error {

	if context == nil {
		return errors.New(ErrNilContext)
	}

	files, _ := ioutil.ReadDir(basePath)
	return f.proccessDirectory(basePath, files, context, excludedFiles)
}

func (f *fileProcessor) proccessDirectory(src string, files []fs.FileInfo, context interface{}, excludedFiles map[string]int) error {
	for _, file := range files {
		filePath := path.Join(src, file.Name())
		_, ok := excludedFiles[filePath]
		if ok {
			renameFiles(file, filePath)
			continue
		}

		if file.IsDir() {
			children, _ := ioutil.ReadDir(filePath)
			if err := f.proccessDirectory(filePath, children, context, excludedFiles); err != nil {
				return err
			}

		} else {
			if err := f.ApplyTemplateToFile(filePath, context); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *fileProcessor) ApplyTemplateToFile(workPath string, templateContext interface{}) error {

	processedContent, err := f.ProcessFile(workPath, templateContext)

	if err != nil {
		return err
	}

	resultingFileName := workPath[:len(workPath)-4]
	resultingFile, err := os.Create(resultingFileName)
	if err != nil {
		return errors.New(ErrCreatingFileFromTemplate + " // " + err.Error())
	}

	if _, err := resultingFile.Write([]byte(processedContent)); err != nil {
		return errors.New(ErrWritingFileFromTemplate + " // " + err.Error())
	}

	os.Remove(workPath)

	return nil
}

func (f *fileProcessor) ProcessFile(fileName string, context interface{}) (string, error) {

	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		return "", errors.New(ErrParseFile + " // " + err.Error())
	}

	return process(tmpl, context)
}

func process(t *template.Template, context interface{}) (string, error) {

	var tmplBytes bytes.Buffer

	ctx := context.(map[string]string)
	if err := t.Execute(&tmplBytes, ctx); err != nil {
		return "", errors.New(ErrTemplateExecute + " // " + err.Error())
	}

	return tmplBytes.String(), nil
}

func renameFiles(file fs.FileInfo, src string) {
	if file.IsDir() {
		files, _ := ioutil.ReadDir(src)
		for _, file := range files {
			filePath := path.Join(src, file.Name())
			renameFiles(file, filePath)
		}
	} else {
		if filepath.Ext(src) == ".tpl" {
			os.Rename(src, src[:len(src)-4])
		}
	}

}
