package main

// >> CODE CRED: https://github.com/squeaky1273 referenced

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type content struct {
	Content string
}

func main() {
	dirPtr := flag.String("dir", "", "dir to look for files in")
	filePtr := flag.String("file", "", "filename to make page from")
	flag.Parse()
	content := readFile(*filePtr)
	fmt.Println(*dirPtr)
	writeTemplateToFile(*filePtr, content)

	fileNames := findFilesInDir(*dirPtr)
	writeTextfileList(fileNames)
}

// reference source for findFilesInDir function:
// https://github.com/TasfiaAddrita/makesite/blob/master/makesite.go
func findFilesInDir(dirName string) []string {
	var textFiles []string

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := strings.Split(file.Name(), ".")
		if len(fileName) > 1 && fileName[1] == "txt" {
			textFiles = append(textFiles, file.Name())
		}
	}
	return textFiles
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeTemplateToFile(filename string, data string) {
	c := content{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	fileName := strings.Split(filename, ".")[0] + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}

}

// write list to stdout
func writeTextfileList(fileNames []string) {
	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}
}
