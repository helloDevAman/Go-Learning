package main

import "fmt"

/*
An interface in Go is a type that defines a set of method signatures.

It acts as a contract: any type whose method set includes all methods
declared in the interface is said to implement that interface.

Interfaces do not store fields/variables—only method signatures.

Implementation is implicit in Go:

A type does not explicitly declare that it implements an interface.
If its method set matches the interface, it automatically satisfies it.

A type must implement ALL methods in an interface to satisfy it.

Syntax:

	type InterfaceName interface {
		Method1(parameters) returnType
		Method2(parameters) returnType
	}

An interface variable stores:

1. The concrete type assigned to it
2. The concrete value of that type

Example:
var s shape = rectangle{length: 2, width: 3}
Here:

Interface Type = shape
Concrete Type = rectangle
Concrete Value = {2, 3}

When calling: s.area()
Go dispatches the call to rectangle.area().
*/
type shape interface {
	area() float64
}

// Value Receiver
type rectangle struct {
	length float64
	width  float64
}

// Value receiver
// Both rectangle and *rectangle satisfy shape.
func (r rectangle) area() float64 {
	return r.length * r.width
}

// METHOD SIGNATURE MUST MATCH EXACTLY
type circle struct {
	radius float64
}

// Does NOT implement shape because signature differs.
// shape requires: area() float64
// circle has: area(float64) float64
func (c circle) areaWithRadius(radius float64) float64 {
	return 3.14 * radius * radius
}

// POINTER RECEIVERS AND INTERFACES
// - If a method has a POINTER receiver:
// - func (r *square) area() float64
// - Then ONLY *square implements the interface, NOT square.
// Because the method belongs to the method set of *square, not square.
type square struct {
	side float64
}

// Pointer receiver
func (s *square) area() float64 {
	return s.side * s.side
}

// A interface with rule pay() for all gateways
type gateway interface {
	pay() error
}

// A struct which will choose which gateway to choose
type checkout struct {
	gw gateway
}

// receiver function implementig checkout struct
func (c checkout) completeCheckout() error {
	err := c.gw.pay()
	if err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}
	fmt.Println("Checkout completed successfully")
	return nil
}

// Razorpay structure
type razorpay struct {
	amount float64
}

// receiver function of struct implementing gateway interface
func (r razorpay) pay() error {
	if r.amount <= 0 {
		return fmt.Errorf("invalid payment amount")
	}
	fmt.Printf("Payment of %.2f completed successfully with Razorpay\n", r.amount)
	return nil
}

type cashfree struct {
	amount float64
}

func (c cashfree) pay() error {
	if c.amount <= 0 {
		return fmt.Errorf("invalid payment amount")
	}
	fmt.Printf("Payment of %.2f completed successfully with Cashfree\n", c.amount)
	return nil
}

func main() {

	var s shape = rectangle{
		length: 2,
		width:  3,
	}

	fmt.Println("Rectangle Area:", s.area()) // Here, the area method of the rectangle struct is

	// COMPILE-TIME ERROR: `circle` does not implement shape because method signature does not match.
	// var s2 shape = circle{radius: 2}

	// INVALID:
	// - square does NOT implement shape because `area()` has pointer receiver
	// var sq shape = square{side: 4}

	// VALID:
	var sq shape = &square{side: 4}
	fmt.Println("Square Area:", sq.area())

	// We can assign any gateway without modifying the actual impl of complete chekout
	// This saves us from violiting `SOLID` principle

	var gw gateway

	gw = razorpay{
		amount: 100,
	}

	// or

	// gw = cashfree{
	// 	amount: 100,
	// }

	// receive function over `checkout` struct
	c := checkout{
		gw: gw,
	}

	if err := c.completeCheckout(); err != nil {
		fmt.Println(err)
	}
}
