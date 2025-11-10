package main

import (
	"fmt"
	"maps"
)

func main() {

	// Decleration only
	var m1 map[string]string
	fmt.Println("Map m1 size is:", len(m1))  // Output: 0
	fmt.Println("Map m1 is nil:", m1 == nil) // Output: true

	// Decleration and Initialization
	m2 := map[int]string{1: "Code"}
	fmt.Println("Map m2 size is:", len(m2))  // Output: 1
	fmt.Println("Map m2 is nil:", m2 == nil) // Output: false

	// Decleration using make function
	m3 := make(map[string]float32, 2)
	fmt.Println("Map m3 size is:", len(m3))  // Output: 0
	fmt.Println("Map m3 is nil:", m3 == nil) // Output: false

	// Write values

	// m1["name"] = "Golang" ---> Output: error
	// This will provide a runtime error because we are trying to add a key-value pair to a nil map,
	// Before adding any we need to initializee the map first like below
	m1 = map[string]string{"camera": "5MP"}
	// or m1 = make(map[string]string)

	m2[2] = "Eat"   // We can assign a key with respectice value like this
	m2[3] = "Sleep" // We can assign a key with respectice value like this
	m3["price"] = 32.90

	// Read values

	fmt.Println("Price is:", m3["price"])
	fmt.Println("First task in the list:", m2[1])

	// Delete values

	delete(m2, 3)
	fmt.Println("Third map after deleting values:", m2)

	// Using values

	value, hasCamera := m1["camera"]

	if hasCamera {
		fmt.Println("Camera is:", value)
	} else {
		fmt.Println("Does not have camera...")
	}

	// Comparing two maps

	phone1 := map[string]interface{}{"battery": 2000, "camera": "20MP"}
	phone2 := map[string]interface{}{"battery": 2000, "camera": "20MP"}

	// isEqual := phone1 == phone2 ---->  Output: error (map can only be comapred to nil)

	isEqual := maps.Equal(phone1, phone2) // With this equal function we can compare same type map as if types are diff then map already diff
	fmt.Println("Is equal:", isEqual)

}
