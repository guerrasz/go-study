package main

import "fmt"

// Overall interface
type Geometry interface {
	area() float64
	perim() float64
}

//& Rectangle struct and implementations
type Rectangle struct {
	wdith, height float64
}

func (r Rectangle) area() float64 {
	return r.wdith * r.height
}

func (r Rectangle) perim() float64 {
	return 2 * (r.wdith + r.height)
}

//& Cricle struct and implementations
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * 3.14159 * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func main() {
	r := Rectangle{
		wdith:  10,
		height: 5,
	}

	c := Circle{
		radius: 5,
	}

	printGeometry(r, "Rectangle")
	printGeometry(c, "Circle", func() { fmt.Printf("Circle diameter: %.2f\n", c.diameter()) })

}

// function accepts any type that implements the Geometry interface
// and any number of functions as extra parameters
func printGeometry(g Geometry, shape string, extras ...func()) {
	fmt.Printf("%s Area: %.2f\n", shape, g.area())
	fmt.Printf("%s Perimeter: %.2f\n", shape, g.perim())
	for _, extra := range extras {
		extra()
	}
	fmt.Println()
}
