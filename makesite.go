package main

// >> CODE CRED: https://github.com/squeaky1273 referenced

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type content struct {
	Content string
}

func main() {
	filePtr := flag.String("file", "", "filename to make page from")
	flag.Parse()
	content := readFile(*filePtr)

	// renderTemplate("template.tmpl", content)
	writeTemplateToFile(*filePtr, content)
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
	err = t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
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
