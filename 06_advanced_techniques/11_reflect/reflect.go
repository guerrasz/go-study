package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	x := 42
	v := reflect.ValueOf(x)

	t := v.Type()

	fmt.Printf("Value %v\n", v)
	fmt.Printf("Type %v\n", t)
	fmt.Printf("Kind %v\n", t.Kind())

	y := 10

	fmt.Printf("Original value of y: %d\n", y)

	v1 := reflect.ValueOf(&y).Elem() // pointer to y
	v1.SetInt(20)                    // set value to 20

	fmt.Printf("New value of y: %d\n", y)

	p := Person{Name: "Alice", Age: 30}
	vp := reflect.ValueOf(p)

	for i := range vp.NumField() {
		fmt.Printf("Field %d, %v\n", i, vp.Field(i))
	}

}
