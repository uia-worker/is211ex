package main

import "fmt"

func main() {
	values := ds2darray.populate()

	// Display first row.
	fmt.Println("Row 1")
	fmt.Println(values[0])

	// Display second row.
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Access an element.
	fmt.Println("First element")
	fmt.Println(values[0][0])

	// Display entire slice.
	fmt.Println("Values")
	fmt.Println(values)
}
