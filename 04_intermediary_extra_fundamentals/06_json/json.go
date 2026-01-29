package main

import (
	"encoding/json"
	"fmt"
)

// struct tags that define the json format
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
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

	// Marshalling
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(jsonPerson))

	jsonDeserialized := `{"name": "Lucas", "age": 22, "email": "lucas@gmail.com", "address": {"city": "São Paulo", "country": "Brazil"}}`

	person2 := Person{}

	err = json.Unmarshal([]byte(jsonDeserialized), &person2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", person2)

	listOfAddress := []Address{
		{
			City:    "São Paulo",
			Country: "Brazil",
		},
		{
			City:    "New York",
			Country: "USA",
		},
		{
			City:    "London",
			Country: "UK",
		},
	}

	fmt.Printf("%v\n", listOfAddress)

	jsonList, err := json.Marshal(listOfAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(jsonList))
}
