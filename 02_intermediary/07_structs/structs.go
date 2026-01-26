package main

import (
	"fmt"
)

// basic person struct
type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	// anonymous field
	Phone
}

type Phone struct {
	cell string
}

// to be nested inside person
type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge(n int) {
	p.age += n
}

func main() {

	lucas := Person{
		firstName: "Lucas",
		lastName:  "Guerra",
		age:       22,
		address: Address{
			city:    "São Paulo",
			country: "Brazil",
		},
		Phone: Phone{
			cell: "123456789",
		},
	}

	lucas.incrementAge(5)

	// when not specified fields are initialized as zero values
	amanda := Person{
		firstName: "Amanda",
		lastName:  "Gois",
	}

	amanda.address.city = "São Paulo"
	amanda.address.country = "Brazil"

	amanda.Phone.cell = "987654321"

	fmt.Printf("First name: %s\nLast name: %s\nAge: %d\n", lucas.firstName, lucas.lastName, lucas.age)
	fmt.Printf("Address: %s, %s\n", lucas.address.city, lucas.address.country)
	//& can be accessed directly because of anonymous field
	fmt.Printf("Cell phone: %s\n", lucas.cell)
	fmt.Println()
	fmt.Printf("First name: %s\nLast name: %s\nAge: %d\n", amanda.firstName, amanda.lastName, amanda.age)
	fmt.Printf("Address: %s, %s\n", amanda.address.city, amanda.address.country)
	//& can be accessed directly because of anonymous field
	fmt.Printf("Cell phone: %s\n", amanda.cell)
	fmt.Println()

	// Anonymous struct
	user := struct {
		username string
		password string
	}{
		username: "lucas",
		password: "123",
	}

	fmt.Printf("Username: %s\nPassword: %s\n", user.username, user.password)

	fmt.Printf("Full name: %s\n", lucas.fullName())

	// structs of the same type are comparable
	fmt.Printf("Lucas and Amanda are equal? %t\n", lucas == amanda)
}
