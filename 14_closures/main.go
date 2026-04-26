package main

import "fmt"

/* A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables;
in this sense the function is "bound" to the variables. */

// closure returns a function that "captures" the local variable `count`.
//
// Even after closure() finishes execution, the returned anonymous function
// still remembers and can modify `count`.
//
// This behavior is called a CLOSURE:
//
// A closure is a function bundled together with the variables
// from its surrounding scope that it references.
func closure() func() int {
	count := 1 // Local variable captured by the returned function
	// or count += 1

	// Anonymous function that modifies and returns the captured variable `count`
	return func() int {
		count++      // Modify captured variable
		return count // Return updated value
	}
}

func main() {
	// Create first independent closure
	increment1 := closure()

	// Create second independent closure
	increment2 := closure()

	// increment1 remembers its own `count`
	fmt.Println("Increment 1:", increment1()) // Output: 2
	fmt.Println("Increment 1:", increment1()) // Output: 3

	// increment2 has its OWN separate `count`
	// It does NOT share state with increment1
	fmt.Println("Increment 2:", increment2()) // Output: 2
	fmt.Println("Increment 2:", increment2()) // Output: 3
}