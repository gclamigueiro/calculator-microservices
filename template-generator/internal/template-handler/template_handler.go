package template_handler

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	ErrParseFile                 = "Error parsing file template"
	ErrNilContext                = "No parameters provided"
	ErrTemplateExecute           = "Error ejecutando template"
	ErrCreatingFileFromTemplate  = "Error creating file from template"
	ErrWritingFileFromTemplate   = "Error writing file from template"
	ErrCreatingDestinationFolder = "Error Creating Destination Folder"
)

// FileProcessor interfaz de procesador de archivos de plantillas
type FileProcessor interface {
	CopyTemplate(src, dest string) error
	ProcessTemplates(basePath string, context interface{}, excludedFiles map[string]int) error
	ApplyTemplateToFile(workPath string, templateContext interface{}) error
}

type fileProcessor struct{}

// NewFileProcessor crea un nuevo template processor
func NewFileProcessor() FileProcessor {
	return &fileProcessor{}
}

func (f *fileProcessor) CopyTemplate(src, dest string) error {

	// First, remove the destination directory in case exists
	os.RemoveAll(dest)

	if err := os.MkdirAll(dest, os.FileMode(0764)); err != nil {
		return errors.New(ErrCreatingDestinationFolder + "->" + err.Error())
	}

	if err := CopyDir(src, dest); err != nil {
		return errors.New(ErrCreatingDestinationFolder + "->" + err.Error())
	}

	return nil
}

func (f *fileProcessor) ProcessTemplates(basePath string, context interface{}, excludedFiles map[string]int) error {

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

	resultingFileName := workPath[:len(workPath)-4]
	resultingFile, err := os.Create(resultingFileName)
	if err != nil {
		return errors.New(ErrCreatingFileFromTemplate + " // " + err.Error())
	}

	ctx := templateContext.(map[string]string)

	// proccessing template
	// processedContent, err := f.ProcessFile(workPath, templateContext)
	tmpl, err := template.ParseFiles(workPath)
	if err != nil {
		return errors.New(ErrParseFile + " // " + err.Error())
	}

	if err := tmpl.Execute(resultingFile, ctx); err != nil {
		return errors.New(ErrTemplateExecute + " // " + err.Error())
	}

	// remove .tpl extension
	os.Remove(workPath)

	return nil
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
