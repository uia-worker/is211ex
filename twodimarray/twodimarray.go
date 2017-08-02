package twodimarray

// Populate a slice of slices (of type integer)
func Populate() [][]int {

	// A table representing customer evaluation
	// Table connects customers to products
	// Many-to-many relationship ?
	product_evaluation := [][]int{}

	// Each Customer evaluates 3 products
	// Input data (usually from outside sources)
	customer1 := []int{2, 0, 3}
	customer2 := []int{5, 2, 0}
	customer3 := []int{3, 3, 1}
	customer4 := []int{0, 2, 2}

	product_evaluation = append(product_evaluation, customer1)
	product_evaluation = append(product_evaluation, customer2)
	product_evaluation = append(product_evaluation, customer3)
	product_evaluation = append(product_evaluation, customer4)

	return product_evaluation

}
