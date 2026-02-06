package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	var pool = sync.Pool{
		New: func() any {
			fmt.Printf("Creating new person\n")
			return &Person{}
		},
	}

	person1 := pool.Get().(*Person)
	person1.Name = "Alice"
	person1.Age = 30

	fmt.Printf("Got person1 name: %s, age: %d\n", person1.Name, person1.Age)

	pool.Put(person1)
	fmt.Printf("Returned person1 to pool\n")

	person2 := pool.Get().(*Person)
	fmt.Printf("Got person2 name: %s, age: %d\n", person2.Name, person2.Age)

	person3 := pool.Get().(*Person)
	person3.Name = "Bob"
	person3.Age = 25
	fmt.Printf("Got person3 name: %s, age: %d\n", person3.Name, person3.Age)

	// return both to pool again
	pool.Put(person2)
	pool.Put(person3)

	fmt.Printf("Returned person2 and person3 to pool\n")

	person4 := pool.Get().(*Person)
	fmt.Printf("Got person4 name: %s, age: %d\n", person4.Name, person4.Age)

	person5 := pool.Get().(*Person)
	fmt.Printf("Got person5 name: %s, age: %d\n", person5.Name, person5.Age)
}
