package main

import "fmt"

// for is the only available loop in go to construct
// but we can achive the similiar while loop settings with for loop

func main() {
	// Basic for loop
	fmt.Println("***Basic for loop***")
	for i := 0; i < 3; i++ {
		fmt.Println("Index is", i)
	}

	// for with while loop similar settings
	fmt.Println("***for with while loop similar settings***")
	i := 0
	for i < 3 {
		fmt.Println("Index is", i)
		i = i + 1
	}

	// for loop with continue that will skip the current iteration
	// In the below example it will skip the code line after continue when the condition is matched for that particular iteration
	fmt.Println("***for loop with continue***")
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Index is", i)
	}

	// for loop with break that will break the execution of for loop after that iteration
	// In the below example the loop iteration will break when condition is matched
	fmt.Println("***for loop with break***")
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Index is", i)
	}

	// for with range
	// The range will provied the number b/w 0 <-> range i.e 11 in the below case
	fmt.Println("***for with range***")
	for i := range 11 {
		fmt.Println("Index is", i)
	}

	// infinite loop
	// fmt.Println("*** infinite loop ***")
	// for {
	// 	fmt.Println("Inddex is N/D")
	// }
}
