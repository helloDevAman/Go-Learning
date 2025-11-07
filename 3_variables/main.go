package main

import "fmt"

func main() {
	// We can declear a variable like below for strings
	var name string = "go"
	var otherName = "Other go"  // Here go will automatically infer the type when assinging a value at the time of decleration
	anotherName := "Another go" // Shorthand decleartion

	fmt.Println(name)
	fmt.Println(otherName)
	fmt.Println(anotherName)

	// We can declear a variable like below for BOOL
	var isAdult bool = true
	var otherIsAdult = false
	anotherIsAdult := true

	fmt.Println(isAdult)
	fmt.Println(otherIsAdult)
	fmt.Println(anotherIsAdult)

	// We can declear a variable like below for int
	var age int = 18
	var othersAge = 12
	anothersAge := 22

	fmt.Println(age)
	fmt.Println(othersAge)
	fmt.Println(anothersAge)

	// We can declear a variable like below for double
	var phLevel float64 = 3.5
	var othersPhLevel = 6.5
	anothersPhLevel := 1.6

	fmt.Println(phLevel)
	fmt.Println(othersPhLevel)
	fmt.Println(anothersPhLevel)

	// Important Notes:

	/**

	# int: It i dafualt integer type and its size depends on the system's architecture.
	  1. On a 32-bit system, an int is 32 bits (4 bytes).
	  2. On a 64-bit system, an int is 64 bits (8 bytes).

	# int32: It is a 32-bit (4-byte) signed integer, regardless of the system's architecture.
	  1. Its value range is from -2,147,483,648 to 2,147,483,647.

	# int64: It is a 64-bit (8-byte) signed integer, regardless of the system's architecture.
	  1. Its value range is from -2,147,483,648 to 2,147,483,647.

	# float32: It is a single-precision floating-point number, occupying 32 bits (4 bytes) of memory.
	  1. It provides about 6–9 decimal digits of precision.

	# float64: It is a double-precision floating-point number, occupying 64 bits (8 bytes).
	  1. It provides about 15–17 decimal digits of precision.

	*/

}
