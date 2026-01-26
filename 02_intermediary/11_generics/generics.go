package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

// pushes an element to the stack
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

// removes last element
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Printf("Stack: ")
	for _, element := range s.elements {
		fmt.Printf("%v ", element)
	}
	fmt.Println()
}

func main() {

	// example of swap with int
	a, b := swap(1, 2)
	fmt.Println(a, b)

	// example of swap with string
	c, d := swap("Hello", "World")
	fmt.Println(c, d)

	// example of swap with float64
	e, f := swap(1.0, 2.0)
	fmt.Println(e, f)

	// example of stack with int
	stack := Stack[int]{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.printAll()
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	// false because stack is empty
	fmt.Println(stack.pop())
}

// swap function that swaps two values of any type
func swap[T any](a, b T) (T, T) {
	return b, a
}
