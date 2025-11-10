package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
		// No need to write the break here GO will automatically terminate excution of next cases if any case is matched
		// If we want execute multiple cases, then we can use multiple condition in one case (see the second example)
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
		// Here default statement is optional
	}

	// siwtch with multiple condition
	switch time.Now().Local().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("boolean")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(30.5)
}
