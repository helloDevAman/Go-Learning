package main

import "fmt"

/*
Structs let us group related data together into one custom type.

They are useful for:
- Modeling real-world entities
- Organizing related fields
- Attaching methods
- Building reusable data structures

Go also supports:
- Constructor functions
- Methods on structs
- Pointer receivers
- Anonymous structs
- Struct embedding (composition)
*/

/*
Address will be embedded inside Employee.
Embedding allows Employee to "contain" address
and access its fields/methods directly.
*/

type address struct {
	city    string
	country string
}

func (a address) fullAddress() string {
	return a.city + ", " + a.country
}

// employee struct with embedded Address struct
type employee struct {
	name    string
	age     int
	id      string
	level   float32
	address // Embedded Struct
}

// Constructor function to create a new employee instance
func newEmployee(name string, age int, id string, city string, country string) *employee {
	return &employee{
		name:  name,
		age:   age,
		id:    id,
		level: 1.0,
		address: address{
			city:    city,
			country: country,
		},
	}
}

// Value receiver method:
// - Works on a COPY of the struct.
// - Does NOT modify original.
// - we are not using a pointer receiver for the struct, so the falsePromote method will not modify
// - the original struct instance because it is working with a copy of the struct.)

func (e employee) falsePromote() {
	e.level += 1.0
}

// Pointer receiver method:
// - Works on original struct.
// - Modifies original value.

func (e *employee) TrulyPromote() {
	e.level += 1.0
}

// We can make getters and setters for the struct fields like this:
// We are not using a pointer receiver for the getName method because it does not modify the struct,
// it simply returns a value. Therefore, a value receiver is sufficient for this method.
func (e employee) getName() string {
	return e.name
}

func main() {
	// Creating an instance of Employee with embedded Address
	employee1 := employee{
		name:  "John Doe",
		age:   30,
		id:    "E12345",
		level: 1.0,
		address: address{
			city:    "Delhi",
			country: "India",
		},
	}

	fmt.Println("Employee 1:", employee1)

	// These fields belong to Address,
	// but are promoted into Employee.
	fmt.Println("Employee City:", employee1.city)
	fmt.Println("Employee Country:", employee1.country)

	// Embedded Methods Are Promoted Too
	fmt.Println("Full Address:", employee1.fullAddress())

	// Creatig a new employee using the constructor function
	employee2 := newEmployee(
		"Jane Smith",
		28,
		"E54321",
		"Mumbai",
		"India",
	)

	fmt.Println("Employee 2 Name:", employee2.getName())
	fmt.Println("Employee 2:", *employee2)

	// Value receiver method does NOT modify original struct
	employee2.falsePromote()
	fmt.Println("After FalsePromote:", *employee2) // Output: Employee Details: {Jane Smith 28 E54321 1.0}
	// Level remains unchanged

	// Pointer receiver method modifies original struct
	employee2.TrulyPromote()
	fmt.Println("After TrulyPromote:", *employee2) // Output: Employee Details: {Jane Smith 28 E54321 2.0}
	// Level increases

	employee3 := &employee{
		name:  "Alice Johnson",
		age:   35,
		id:    "E67890",
		level: 1.0,

		// Using struct literal to initialize the embedded struct directly
		address: address{
			city:    "Bangalore",
			country: "India",
		},
	}

	fmt.Println("Employee 3:", *employee3)

	// Useful for temporary one-off structured data.
	language := struct {
		Name    string
		Version string
	}{
		Name:    "Go",
		Version: "1.20",
	}

	fmt.Println("Language Details:", language)
}
