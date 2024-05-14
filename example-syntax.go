package main

import "fmt"

func example() {
	// Declare and initialize variables
	var age int = 10
	name := "John"

	// Print the variables
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)

	// If statement example
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Create an array of numbers
	numbers := []int{1, 2, 3, 4, 5}

	// Loop over the numbers
	for i, number := range numbers {
		fmt.Printf("Number %d: %d\n", i, number)
	}

	// Call a function
	result := add(10, 20)
	fmt.Printf("Result: %d\n", result)

}

// Creating a function
func add(a, b int) int {
	return a + b
}
