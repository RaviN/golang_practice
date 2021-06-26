package main

import "fmt"

const (
	first = iota + 6
	second
	pi = 3.1415
)

const (
	third = iota
	fourth
)

func main() {
	/*
		Value Types
	*/
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Ravi"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	/*
		Pointers
	*/
	var name *string = new(string)

	*name = "Ravi N"

	fmt.Println(*name)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Bob"
	fmt.Println(ptr, *ptr)

	/*
		Constants
	*/
	const c1 = 3
	fmt.Println(c1 + 3)

	fmt.Println(c1 + 1.2)

	// iota
	// 6 7 0 1
	fmt.Println(first, second, third, fourth)
}
