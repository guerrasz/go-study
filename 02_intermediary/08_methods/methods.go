package main

import "fmt"

type Rectangle struct {
	lenght float64
	width  float64
}

// method with value receiver
func (rectangle Rectangle) area() float64 {
	return rectangle.lenght * rectangle.width
}

// method with pointer receiver
func (rectangle *Rectangle) scale(factor float64) {
	rectangle.lenght *= factor
	rectangle.width *= factor
}

type MyInt int

// method on a non-struct type
func (i MyInt) isPositive() bool {
	return i > 0
}

func main() {
	rect := Rectangle{
		lenght: 10,
		width:  5,
	}

	fmt.Printf("Area: %.2f\n", rect.area())

	rect.scale(2)
	fmt.Printf("Area after scaling dimensions by 2: %.2f\n", rect.area())

	// methods on non-struct types
	i := MyInt(5)
	fmt.Printf("Is %d positive? %t\n", i, i.isPositive())
	i = MyInt(-3)
	fmt.Printf("Is %d positive? %t\n", i, i.isPositive())
}
