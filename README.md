# is211ex
Some examples while evaluating a complaint in IS-211 summer 2017

Choosing data structure
Choice between a two dimensional array and a map structures
Map structure is more appropriate if data is sparse (many empty cells in a 2D table representation)
In Golang a 2D array can be made as
(1) slice of slices [][]int{}
(2) 2d array [2][2]int{}
