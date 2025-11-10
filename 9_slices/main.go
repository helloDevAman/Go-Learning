package main

import "fmt"

/**
1. Slice -> dynamic array
2. Many more predefined methods as compared to arrays
*/

func main() {
	// Decleration (no initialization)
	var nums []int           // Declares a slice of int, but does not initialize it.
	fmt.Println(nums)        // Output:[] (empty slice i.e nil)
	fmt.Println(nums == nil) // Output: ture (As nums not initialized)
	fmt.Println(len(nums))   // Output: 0
	fmt.Println(cap(nums))   // Output: 0

	nums = append(nums, 1) // We can append this
	fmt.Println("Nums value after append:", nums)

	// Decleration and Initialization
	var values = []bool{} // Declares and initializes nums as an empty, non-nil slice.
	fmt.Println("Values are:", values)
	fmt.Println("Values length:", len(values))
	fmt.Println("Values slices is nil:", values == nil) // Output: false (As values is an empty, non-nil slice)
	fmt.Println("Values capacity:", cap(values))        // Output: 0

	// Decleration using make function

	/**
	make() takes three params:
	1 -> type
	2 -> size
	3 -> capacity (optional, defaults to size if not provided)
	*/
	var names = make([]string, 0) // Empty but allocated like []int{}, so it’s empty but non-nil
	fmt.Println(names)            // Output: [] (empty but non-nil)
	fmt.Println(cap(names))       // Output: 0 (empty but non-nil)
	fmt.Println(names == nil)     // Output: false

	// Slice with auto capacity increase
	auto := make([]int, 0, 1)
	fmt.Println("Initial slice:", auto)
	fmt.Println("Initial length:", len(auto))
	fmt.Println("Initial cap:", cap(auto))

	for i := range 11 {
		auto = append(auto, i)
		fmt.Println("Initial slice:", auto)
		fmt.Println("Initial length:", len(auto))
		fmt.Println("Initial cap:", cap(auto))
	}

	/**
	Note:
	- When we append an element beyond the current capacity, Go automatically increases the capacity.
	- The growth pattern is implementation-dependent, but generally doubles each time (1, 2, 4, 8, 16...).
	- This makes slices efficient for dynamic array usage.
	*/
}
