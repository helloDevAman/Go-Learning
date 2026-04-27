package main

import "fmt"

/*

In Go we can define our own types as well i.e user defined types

Syntax:
type NewType existingType (string, int, func, etc.)

Example:
type Age int
type Username string
type Callback func()

This creates a NEW distinct type

IMPORTANT:
Even if the underlying type is the same,
Go treats the custom type as a separate type.
 - Better readability / domain meaning
 - Type safety
 - Attach methods to built-in-like values
 - Create reusable function signatures

*/

// Custom types are distinct types
type Age int
type Username string

func printAge(a Age) {
	fmt.Println("Age:", a)
}

// We can add receiver methods to a custom type too
func (a Age) isAdult() bool {
	return a >= 18
}

// We can have custom types as `func()` too
type VoidCallback func()

type MathOperation func(int, int) int

func execute(vc VoidCallback) {
	vc()
}

func calculate(a int, b int, op MathOperation) int {
	return op(a, b)
}

// Custom types ca also implement interfaces

type printer interface {
	print()
}

func (u Username) print() {
	fmt.Println("Username:", u)
}

func main() {

	var age Age = 21
	var user Username = "Rahul"

	printAge(age)

	fmt.Println("Is Adult:", age.isAdult())

	var x int = 25

	// INVALID:
	// printAge(x)

	// VALID:
	printAge(Age(x))

	var cb VoidCallback = func() {
		fmt.Println("Callback executed!")
	}

	execute(cb)

	add := func(a int, b int) int {
		return a + b
	}

	result := calculate(10, 20, add)
	fmt.Println("Addition Result:", result)

	var p printer = user
	p.print()
}
