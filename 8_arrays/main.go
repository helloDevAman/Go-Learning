package main

import "fmt"

func main() {
	// Numbered sequence of fixed length (Decleared only)
	var nums [4]int
	fmt.Println(len(nums)) // We can get array length with buit in len()

	/**
	// In case we have not assigned the values in arrays
	// Deafult(zeroed) value will be assigned at the time of decleration
	// i.e int -> 0, bool -> false, string -> "", flaot32/float64 -> 0.0
	*/
	fmt.Println(nums) // Output: [0 0 0 0]

	nums[0] = 1          // We can assign values to specific index
	fmt.Println(nums[0]) // Access specific index values

	// Decleared and initialized
	var names = [2]string{"Aman", "John"} // If we will assign the more values than the array len decleared it will give an error saying [index out of bound]
	// names := [2]string{"Aman", "John"} // We can decleare like this too
	fmt.Println("Name array is:", names)

	// 2D-Array
	vals := [2][3]int{{2, 3, 4}, {23, 34, 43}}
	fmt.Println("2D Araay:", vals)
}

// Note: When to use array
// -> When we have fixed size, len is predictable
// -> When  we have memory constraints
// -> Constant time access
