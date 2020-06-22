package main

// <<<<<<< HEAD
// unknown purpose of this string. it was in the starter code
// >> CODE CRED: https://github.com/squeaky1273
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
	content := readFile()
//     # command-line-arguments
// ./makesite.go:20:21: too many arguments in call to readFile
// 	have (string) :: removed *filePtr from ()
// 	want ()

	renderTemplate("template.tmpl", content)
	writeTemplateToFile("template.tmpl", *filePtr)
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}


// >>>>>>> 9514ac8a2c135a448a2b15a4b246dcd5d59ee7bf
// unknown what the purpose is of this string. it was in the starter code
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
