package main

import "fmt"

/*
A pointer stores the memory address of another variable.

Why use pointers?
- To modify original values inside functions
- To avoid copying large data structures
- To share access to the same data

Syntax:
    var p *int   // pointer to an int
    p = &x       // store address of x in p

Operators:
    &  -> "address of" operator
    *  -> "dereference" operator
*/


// By Value:
// - A COPY of the variable is sent to the function
// - Changes inside the function do NOT affect the original

func byValue(count int) {
	count = count + 1
	fmt.Println("Inside byValue():", count)
}

// By reference (Pointters):
// - A pointer to the variable is sent to the function
// - Changes inside the function affect the original variable
func byReference(count *int) {
	*count = *count + 1
	fmt.Println("Inside byReference():", *count)
}

func main() {
	count := 10
	fmt.Println("Initial count:", count)

	byValue(count)
	fmt.Println("After byValue():", count) // Output: 10, original value unchanged

	byReference(&count)
	fmt.Println("After byReference():", count) // Output: 11, original value modified

	// Store the memory address of count in a pointer variable
	var p *int
	p = &count
	fmt.Println("Memory address of count:", p)
	

	// Chnage the value of count using the pointer
	*p = *p + 1
	fmt.Println("After modifying count through pointer:", count) // Output: 12
}
