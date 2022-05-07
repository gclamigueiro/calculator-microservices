package main

/*
Small utility to add the .tpl extension to every file in a folder.
If the file extension is already .tpl, it will not be modified.
*/

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func main() {
	src := "./templates/go-kit"
	files, _ := ioutil.ReadDir(src)
	renameFiles(src, files)
}

func renameFiles(src string, files []fs.FileInfo) {
	for _, file := range files {
		filePath := path.Join(src, file.Name())
		if file.IsDir() {
			children, _ := ioutil.ReadDir(filePath)
			renameFiles(filePath, children)
		} else {
			if filepath.Ext(filePath) != ".tpl" {
				fmt.Println(filePath)
				os.Rename(filePath, filePath+".tpl")
			}
		}
	}

}
