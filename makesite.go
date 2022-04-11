package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {

	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	//fmt.Print(string(fileContents))

	page := Page{
		TextFilePath: "",
		TextFileName: "new",
		HTMLPagePath: "new.html",
		Content:      string(fileContents),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	file := flag.String("file", "defaultValue", " Help text.")
	flag.Parse()
	fmt.Print(*file)

	// Create a new, blank HTML file.
	newFile, err := os.Create(*file)
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)
}
