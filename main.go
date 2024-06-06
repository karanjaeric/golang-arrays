package main

import "fmt"

func main() {
	var x [3]int
	var index = 2
	//writing and reading to an array
	fmt.Printf("Variable x of three ints %v\n", x)
	fmt.Printf("index %v value is  %v\n", index, x[index])
	//writing and reading to an array
	x[2] = 10
	fmt.Printf("index %v value is  %v\n", index, x[index])
	//literak array declaration
	var y = [4]int{20, 40, 30, 85}
	fmt.Printf("Array Y value is  %v\n", y)
	//using ellipse in a literal array
	var ellipseArray = [...]int{20, 40, 30, 85}
	fmt.Printf("Ellipse Array value is  %v\n", ellipseArray)
	//using == operator
	fmt.Printf("Is Y and Ellipse arrays equal?  %v\n", y == ellipseArray)
	//using != operator
	fmt.Printf("Is Y array NOT EQUAL to Ellipse arrays equal?  %v\n", y != ellipseArray)

	//finding the length of an array using len function
	fmt.Printf("Ellipse Array is of size  %v\n", len(ellipseArray))

}
