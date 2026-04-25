package main

import "fmt"

// Imagine if we can pass any number of arguments to a function?
// In Go, you can do that with variadic functions.
// A variadic function is a function that can take a variable number of arguments.
func main() {
	// e.g `fmt.Println` is a pre-defined variadic function that can take any number of arguments of any type.
	fmt.Println(1, 2, 3, 4, "Goat")

	// Variadic function that takes a variable number of integers and returns their sum.
	fmt.Println(sum(1, 2, 3, 4, 5))

	// We can also pass a slice of integers to the variadic function by using the ellipsis (...) after the slice variable.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))

	// We can also pass a array of integers to the variadic function by using the ellipsis (...) after the array variable.
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:]...))

	// Variadic function that can take any type of arguments.
	anything(1, "Hello", 3.14, true)

	// Variadic function with a non-variadic parameter.
	nonVariadicParameter("Alice", 1, 2, 3, 4, 5)
}

// We can define a variadic function by using the ellipsis (...) before the type of the last parameter.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// We can also use the empty interface `interface{}` to accept any type of arguments.
// from v1.18 we can use `any` instead of `interface{}` as it is an alias for the empty interface.
func anything(nums ...interface{}) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

// Variadic function with a non-variadic parameter.
func nonVariadicParameter(name string, nums ...int) {
	fmt.Printf("Hello, %s! The sum of the numbers is: %d\n", name, sum(nums...))
}

// Note: variadic arguments are packaged into a slice automatically.
// Note: a variadic function can only have one variadic parameter and it must be the last parameter in the function signature.
