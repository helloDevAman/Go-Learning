/**
1.	Entry Point: When we use the go run or go build command, the Go toolchain looks for the source file(s) that declare package main.
2. 	The main Function: Within the package main, the compiler then specifically searches for the mandatory func main() function.
3. 	This function serves as the entry point—the very first line of code that gets executed when the compiled program starts running.
4. 	If file has package main and func main(), the go build command produces a single, runnable executable file (a standalone program).
5. 	Platform Dependent: The Go toolchain is smart; when you run go build without specifying a target OS, it builds an executable designed to run on the current OS of your machine.
	It automatically adds the .exe extension only when the target is Windows.
6.	Cross-Compilation: One of Go's most powerful features is cross-compilation. You can build a Windows executable (.exe) even while sitting on a Mac or Linux machine using environment variables:
	[GOOS=windows go build], This command, run on any OS, will produce an executable file ending in .exe.
7.	If a package is not named main (e.g., package utilities), it is treated as a library or module. A library cannot be executed standalone;
	it can only be imported and used by another program that does contain package main.
*/

package main // Go will look for this main package when we are running the the program standalone

import "fmt" // fmt is the library in which formatting and printing methods are there it stands for formatting

func main() {

	// In go we can print a string with below command
	fmt.Println("Hello World!")

	// In go we can print a number like below
	fmt.Println(1)

	// In go we can concatenate two strings like below and print
	fmt.Println("Hello" + " " + "World!")

	// In go we can concatenate a string and an integer
	fmt.Println("My lucky number is", 7)

	// In go we can concatenate a string and a floating-point number
	fmt.Println("The value of pi is approximately", 3.14)

	// In go we can concatenate a string and a bool
	fmt.Print("Access granted is: ", false)
}
