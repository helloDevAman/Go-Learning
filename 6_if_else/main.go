package main

import "fmt"

func main() {

	// Normal if-else cases
	age := 11

	if age < 12 {
		fmt.Println("No, person is a kid")
	} else if age < 18 {
		fmt.Println("No, person is a teen")
	} else {
		fmt.Println("Yes, prsona is an adult")
	}

	// if-else with logical operator

	var role = "admin"
	var isAllowed = false

	if role == "admin" && isAllowed {
		fmt.Println("User is admin and allowed to modify...")
	} else if role == "admin" || isAllowed {
		fmt.Println("User is an admin or allowed to modify...")
	} else {
		fmt.Println("User not an admin and not allowed to modify too...")
	}

	// if-else with inline variable decleration

	if percentage := 70; percentage < 40 {
		fmt.Println("Student has secured C garde")
	} else if percentage < 50 {
		fmt.Println("Student has secured B grade")
	} else if percentage < 60 {
		fmt.Println("Student has secured A grade")
	} else {
		fmt.Println("Student has secured A+ grade")
	}
}
