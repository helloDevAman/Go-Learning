package main

import "fmt"

// Add two numbers and return the result
func add(a int, b int) int {
	return a + b
}

// If consecutive parameters have the same type,
// we can omit the type for all but the last parameter in that group.
func subtract(a, b float32, c, d, e int) int {
	return int(a-b) + c - d + e
}

// We can also return multiple values (same or different types) from a function.
// The return types are defined in the function signature
// as a comma-separated list enclosed in parentheses.
func multipleReturnValues() (string, string) {
	return "Delhi", "Mumbai"
}

// Order of returned values must match the return types
// defined in the function signature.
func multipleReturnValuesWithOrder(x, y string, z int) (string, string, int) {
	return y, x, z
}

// We can also assign a function to a variable and call it using that variable.
// This is possible because in Go, functions are first-class citizens.
func processTask(task func(who string) (string, string, int)) (firstname, lastname string, age int) {
	return task("Bob")
}

// functions can also return functions
func returnFunction() func(name string) string {
	return func(name string) string {
		return "Hello, " + name + "! This isfrom the returned function!"
	}
}

func main() {
	// Calling function and printing the result with two parameters
	sum := add(5, 3)
	fmt.Println("Sum of 5 and 3 is:", sum)

	// Calling function and printing the result with multiple parameters of different types
	difference := subtract(10.5, 4.2, 3, 2, 1)
	fmt.Println("Result:", difference)

	// Calling function and printing the result with multiple return values
	fmt.Println(multipleReturnValues())

	// Calling function and printing the result with multiple return values in a specific order
	v1, v2, v3 := multipleReturnValuesWithOrder("Hello", "World", 42)
	fmt.Printf("v1: %s, v2: %s, v3: %d\n", v1, v2, v3)

	// Assigning a function to a variable and calling it
	fun := func(who string) (string, string, int) {
		if who == "Alice" {
			return "Alice", "Smith", 25
		}
		if who == "Bob" {
			return "Bob", "Johnson", 35
		}
		return "John", "Doe", 30
	}
	// Calling the processTask function with the assigned function and printing the result
	value1, value2, value3 := processTask(fun)
	fmt.Printf("value1: %s, value2: %s, value3: %d\n", value1, value2, value3)

	// Calling the function that returns a function and then calling the returned function
	function := returnFunction()
	message := function("Alice")
	fmt.Println(message)
}
