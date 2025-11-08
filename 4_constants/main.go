package main

import "fmt"

func main() {
	// We can decleare constants variables like below
	const name = "Aman"
	const age = 26

	// Below operation (reassigning) is not supported on a constants
	//  name = "Aman Kumar"
	//  age = 28

	// We can also decleare multiple costants together
	const (
		address = "Muzaffarpur"
		mobile  = "+9189797"
	)

	// Shorthand declearation is not supported like below
	// email := "helloaman"

	fmt.Print("My name is ", name, ", age is ", age, ", address is ", address, " and email is ", email)
}

// We can also decleare the constants variable outside the functions
const email = "hello@joho.com"
