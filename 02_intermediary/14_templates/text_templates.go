package main

import (
	"os"
	"text/template"
)

func main() {
	//^ more verbose way to create a template
	//tmpl := template.New("name")
	// parse template based on string .Name is a variable to be reused
	//tmpl, err := tmpl.Parse("Hello {{.Name}}\n")
	//! error handling
	// if err != nil {
	// 	panic(err)
	// }

	// concise way less verbose auto handling errors
	tmpl := template.Must(template.New("name").Parse("Hello {{.Name}}\n"))

	// execute template map of string -> string, the key is the template var and value is the value to be replaced
	err := tmpl.Execute(os.Stdout, map[string]string{"Name": "Guerra"})
	if err != nil {
		panic(err)
	}
}
