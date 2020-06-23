package main

// >> CODE CRED: https://github.com/squeaky1273 referenced

import (
	"io/ioutil"
    "html/template"
    "flag"
    "os"
    "strings"
)

type content struct {
	Content string
}

func main() {
    filePtr := flag.String("file", "", "filename")
	flag.Parse()
	content := readFile(*filePtr)
//     # command-line-arguments
// ./makesite.go:20:21: too many arguments in call to readFile
// 	have (string) :: removed *filePtr from ()
// 	want ()

	renderTemplate("template.tmpl", content)
	writeTemplateToFile("template.tmpl", *filePtr)
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}


func renderTemplate(filename string, data string) {
	c := content{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , c)
	if err != nil {
		panic(err)
	}
}

func writeTemplateToFile(filename string, data string) {
	c := content{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

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
