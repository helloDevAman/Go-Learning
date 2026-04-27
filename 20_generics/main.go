package main

import "fmt"

/*

Generics allow writing reusable code that works with multiple types while preserving type safety.
Introduced in Go 1.18.


Syntax:

func Name[T any](param T) { ... }
T is a type parameter. 'any' means T can be any type.

*/

// Without generic need seprate fucntions for each type `printerArray`, `printerSlice` and `printerStringSlice`
func printerArray(s [3]int) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func printerSlice(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func printerStringSlice(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// With generic one function works for multiple slice types.
func printer[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// struct `box` with generic type
type box[T any] struct {
	value T
}

// struct `pair` with generic type
type pair[K any, V any] struct {
	key   K
	value V
}

func main() {
	fmt.Println("::::::Generics:::::")

	// In case of `printer` accept array of int
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Array without generic...")
	printerArray(arr)

	// In case of `printer` accepts slice of int
	var s []int
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println("Slice without generic...")
	printerSlice(s)

	// In case of `printerStringSlice` accepts slice of string
	var str []string
	str = append(str, "one")
	str = append(str, "two")
	str = append(str, "three")
	fmt.Println("Slice of string without generic...")
	printerStringSlice(str)

	// In the above three examples we have seen we need to create different methods for each type
	// Here genreics comes in picture created a method `printer` with `T`

	var arr2 [3]int
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	fmt.Println("Array with generic...")
	printer(arr2[:])

	var s2 []int
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	fmt.Println("Slice with generic...")
	printer(s2)

	var str2 []string
	str2 = append(str2, "one")
	str2 = append(str2, "two")
	str2 = append(str2, "three")
	fmt.Println("Slice with string with generic...")
	printer(str2)

	// Using generic structs
	myBox := box[int]{value: 100}
	fmt.Println(myBox.value)

	myPair := pair[string, int]{key: "key", value: 100}
	fmt.Println(myPair.key, myPair.value)
}
