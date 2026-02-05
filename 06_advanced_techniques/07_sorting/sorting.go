package main

import (
	"fmt"
	"sort"
)

// base struct
type Person struct {
	Name string
	Age  int
}

// by age type
type ByAge []Person

// common interface implementations
func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// by name type
type ByName []Person

// common interface implementations
func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

// workaround to avoid code duplication for common interface implementations by using a function type that takes a comparison function
type By func(p1, p2 *Person) bool

// PersonSorter struct that implements sort.Interface and uses the By function type to compare elements
type PersonSorter struct {
	people []Person
	by     By
}

// swap and len implementations that use the people slice don't need custom comparison logic so they can be shared
func (ps *PersonSorter) Len() int {
	return len(ps.people)
}

func (ps *PersonSorter) Swap(i, j int) {
	ps.people[i], ps.people[j] = ps.people[j], ps.people[i]
}

// less implementation that uses the provided comparison function to compare elements at index i and j
func (ps *PersonSorter) Less(i, j int) bool {
	return ps.by(&ps.people[i], &ps.people[j])
}

// Sort method to sort a slice of Person using the provided By comparison function
func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	// Base Array
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Dave", 20},
		{"Eve", 28},
	}

	// NOTE: sort by age
	fmt.Printf("Unsorted array %v\n", people)

	sort.Sort(ByAge(people))

	fmt.Printf("Sorted arrray %v\n", people)

	// NOTE: sort by name
	sort.Sort(ByName(people))

	fmt.Printf("Sorted arrray %v\n", people)

	// NOTE: using the PersonSorter with a custom comparison function to sort by age
	// define the comparison function for age
	ageSorter := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}

	// use the PersonSorter to sort by age converting the function to By type	
	By(ageSorter).Sort(people)

	fmt.Printf("Sorted by age using PersonSorter %v\n", people)
}
