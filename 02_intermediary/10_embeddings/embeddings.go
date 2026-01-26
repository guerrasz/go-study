package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person     // Embedded struct
	employeeId int
	salary     float64
}

func (p Person) introduce() {
	fmt.Printf("Hi, my name is %s and I'm %d years old\n", p.name, p.age)
}

// override method with embedding
func (e Employee) introduce() {
	fmt.Printf("Hi, my name is %s and my ID is %d\n", e.name, e.employeeId)
}

func main() {
	p := Person{
		name: "Lucas",
		age:  22,
	}

	e := Employee{
		Person:     p,
		employeeId: 123,
		salary:     1000,
	}

	p.introduce()
	e.introduce()

	// access to embedded fields
	fmt.Printf("Employee name: %s\n", e.name)
	fmt.Printf("Employee age: %d\n", e.age)
	fmt.Printf("Employee ID: %d\n", e.employeeId)
	fmt.Printf("Employee salary: %.2f\n", e.salary)
}
