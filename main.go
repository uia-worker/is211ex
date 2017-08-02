package main

import "fmt"
import "./twodimarray"

func main() {
	values := twodimarray.Populate()
	// Display first row.
	fmt.Println("Customer #1 evaluates Product #1-3")
	fmt.Println(values[0])

	// Display second row.
	fmt.Println("Customer #2 evaluates Product #1-3")
	fmt.Println(values[1])

	// Access an element.
	fmt.Println("Customer #1 evaluates Product #1")
	fmt.Println(values[0][0])

	// Display entire slice.
	fmt.Println("Customer #1-4 evaluates Product #1-3")
	fmt.Println(values)
}
