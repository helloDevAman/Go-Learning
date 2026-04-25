package main

import "fmt"

// We use range ti iterate over data structures like arrays, slices, maps, strings, and channels.
// It returns one or two values depending on the data structure being iterated over.
// The first value is the index or key, and the second value is the corresponding element or value.
func main() {
	// range with strings
	fmt.Println("range with strings:")
	str := "Hello, World!"
	for index, char := range str {
		fmt.Println(index, string(char)) // Here we convert the rune to a string for better readability, otherwise it would print the Unicode code point of the character.
	}

	// Note: When iterating over a string with range,
	// it iterates over Unicode code points (runes), not raw bytes.
	// The index returned by range is the byte position of the rune
	// in the UTF-8 encoded string, not the character position.
	// Multi-byte characters (such as emojis or many non-Latin characters)
	// occupy more than one byte in UTF-8.
	// e.g:  s := "Hello😊Dev"
	// Iteration yields:
	//   index 5  -> '😊'  (emoji occupies 4 bytes)
	//   index 9  -> 'D'
	// The index jumps because the emoji takes multiple bytes in UTF-8.
	// range with arrays

	fmt.Println("range with arrays:")
	arr := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// range with slices
	fmt.Println("range with slices:")
	slice := [][]float32{{23.0, 43.4}, {12.5, 34.6}, {45.7, 67.8}}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// range with maps
	fmt.Println("range with maps:")
	m := map[string]string{"name": "Alice", "age": "30", "city": "New York"}
	for key, value := range m {
		fmt.Println(key, value) // The order of iteration over a map is not guaranteed, so the output may vary each time you run the program.
	}

	// or if you need only keys for maps
	fmt.Println("range with maps (only keys):")
	for key := range m {
		fmt.Println(key)
	}

	// Similarly, if you need only values for maps
	fmt.Println("range with maps (only values):")
	for _, value := range m {
		fmt.Println(value)
	}

	// Note: if index is not needed for arrays or slices keep in mind that the first parameter is the index/key.
	// So, if you want to ignore the index, you can use the blank identifier (_) to discard it.
	// For example, if you want to iterate over a slice and only care about the values, you can do:
	fmt.Println("range with slices (only values):")
	for _, value := range slice {
		fmt.Println(value)
	}
}
