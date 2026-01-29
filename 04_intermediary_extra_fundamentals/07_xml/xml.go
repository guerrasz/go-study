package main

import (
	"encoding/xml"
	"fmt"
)

// Person struct with XML tags
// xml.Name field allows specifying the root element name
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
	Address Address  `xml:"address"`
}

type Address struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

// People struct to handle lists of Person
type People struct {
	XMLName xml.Name `xml:"people"`
	Persons []Person `xml:"person"`
}

func main() {
	person := Person{
		Name:  "Lucas",
		Age:   22,
		Email: "lucas@gmail.com",
		Address: Address{
			City:    "São Paulo",
			Country: "Brazil",
		},
	}

	// Marshalling (Struct to XML)
	fmt.Println("--- Marshalling ---")
	xmlPerson, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n\n", string(xmlPerson))

	// Unmarshalling (XML to Struct)
	fmt.Println("--- Unmarshalling ---")
	xmlInput := `
	<person>
		<name>Lucas</name>
		<age>22</age>
		<email>lucas@gmail.com</email>
		<address>
			<city>São Paulo</city>
			<country>Brazil</country>
		</address>
	</person>`

	var person2 Person
	err = xml.Unmarshal([]byte(xmlInput), &person2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized: %+v\n\n", person2)

	// Handling Lists
	fmt.Println("--- List of Items ---")
	people := People{
		Persons: []Person{
			{
				Name: "John",
				Age:  30,
				Address: Address{
					City:    "New York",
					Country: "USA",
				},
			},
			{
				Name: "Alice",
				Age:  25,
				Address: Address{
					City:    "London",
					Country: "UK",
				},
			},
		},
	}

	xmlList, err := xml.MarshalIndent(people, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(xmlList))
}
