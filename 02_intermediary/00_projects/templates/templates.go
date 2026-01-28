package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// new buffer reading from stdin
	reader := bufio.NewReader(os.Stdin)

	// welcome message
	fmt.Printf("Enter your name: ")

	// reads until \n character
	name, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	// trim space from input
	name = strings.TrimSpace(name)

	// define named templates
	templates := map[string]string{
		"welcome":      "Hello {{.name}}!",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Oops! An error ocurred {{.errorMessage}}",
	}

	// parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	// for as while
	for {
		// show menu
		fmt.Println("\nMenu")
		fmt.Println("1. Welcome")
		fmt.Println("2. Notification")
		fmt.Println("3. Error")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		fmt.Println()

		switch choice {
		case "1":
			// define template
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{"name": name}

		case "2":
			// define template
			tmpl = parsedTemplates["notification"]
			fmt.Print("Enter your notification: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			data = map[string]any{"name": name, "notification": notification}
			fmt.Println()

		case "3":
			// define template
			tmpl = parsedTemplates["error"]
			fmt.Print("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			data = map[string]any{"errorMessage": errorMessage}
			fmt.Println()

		case "4":
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid option")
			continue
		}

		// execute template
		err = tmpl.Execute(os.Stdout, data)
		fmt.Println()
		if err != nil {
			panic(err)
		}
	}
}
