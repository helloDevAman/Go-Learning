package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Goroutine: A goroutine is a lightweight thread managed by the Go runtime.
Use the `go` keyword to run a function concurrently.

Example:
	go someFunction()z

Concurrency:
	Multiple tasks make progress together.

Parallelism:
	Multiple tasks execute at the exact same time.

Go supports both.
*/




/*
Normal function calls are blocking.
Each task waits for the previous task to finish.

Expected Output:
	Doing task: 0
	Doing task: 1
	Doing task: 2
	...
*/
func blockingExample() {
	fmt.Println("* Blocking Execution *")

	for i := 0; i < 5; i++ {
		task(i)
	}
}

/*
Adding `go` before a function call creates a goroutine.

Important:
	The output order is NOT guaranteed.

Why?
	Because the Go scheduler decides which goroutine runs first.

Possible Output:
	Doing task: 3
	Doing task: 1
	Doing task: 4
	Doing task: 0
	Doing task: 2

*/

func goroutineExample() {
	fmt.Println("\n* Goroutine Execution *")

	for i := 0; i < 5; i++ {
		go task(i)
	}

	// IMPORTANT ISSUE:
	// If we don't pause or wait,
	// the main function exits immediately.
	// Then all goroutines terminate before completion.

	// time.Sleep() is NOT the ideal solution.
	// It is only used for demo/testing purposes.

	time.Sleep(time.Second)
}

/*
PROBLEM:
	Goroutine captures the loop variable by reference.
Wrong Example:
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

Possible Output:
	5
	5
	5
	5
	5

Why?
	By the time goroutines execute,
	the loop has already completed.
*/

func wrongClosureExample() {
	fmt.Println("\n* Wrong Closure Example *")

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Wrong closure value:", i)
		}()
	}

	time.Sleep(time.Second)
}

/*
SOLUTION:
	Pass the loop variable as a parameter.

Correct Example:
	go func(i int) {
		fmt.Println(i)
	}(i)

Now every goroutine gets its own copy.
*/

func correctClosureExample() {
	fmt.Println("\n* Correct Closure Example *")

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("Correct closure value:", i)
		}(i)
	}

	time.Sleep(time.Second)
}

/*
Instead of using time.Sleep(),
	use sync.WaitGroup.

Why?
	Because WaitGroup properly waits for goroutines.

Functions:
	wg.Add(1)   -> increase counter
	wg.Done()   -> decrease counter
	wg.Wait()   -> wait until counter becomes zero
*/

func waitGroupExample() {
	fmt.Println("\n* WaitGroup Example *")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		// Increase counter
		wg.Add(1)

		go func(id int) {

			// Decrease counter after completion
			defer wg.Done()

			task(id)

		}(i)
	}

	// Wait for all goroutines
	wg.Wait()

	fmt.Println("All goroutines completed")
}

func task(id int) {
	fmt.Println("Doing task:", id)
}

func main() {

	blockingExample()

	goroutineExample()

	wrongClosureExample()

	correctClosureExample()

	waitGroupExample()
}

/*
1. Goroutines are lightweight
   - Small memory footprint
   - Managed by Go runtime

2. Avoid excessive goroutines
   - Can consume huge memory

3. Prefer WaitGroup over Sleep

4. Use Mutex for shared memory safety

5. Use Channels for communication

6. Detect race conditions:
	go run -race main.go

7. Common beginner mistakes:
   - Forgetting wg.Done()
   - Using Sleep for synchronization
   - Shared variable race conditions
   - Closure capture bugs
   - Not closing channels
*/