package main

import "fmt"

/*

Go does NOT have a native enum keyword.
Instead, enum-like behavior is commonly created using:

Instead, enum-like behavior is commonly created using:

	1. A custom named type
	2. A group of constants
	3. Optionally iota for auto-incrementing values

Enums in Go are therefore convention-based, not a special language feature.

Syntex:

type UserDefinedType int

const (
	EnumValue1 UserDefinedType = iota
	EnumValue2
	EnumValue3
)


Example:

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

How iota works
iota starts at 0 inside a const block and increments by 1 for each new line.
Equivalent to:

const (
	Pending Status = 0
	Approved Status = 1
	Rejected Status = 2
)
*/

// enum of Exchange with string typed values
type Exchange string

const (
	NSE Exchange = "nse"
	BSE Exchange = "bse"
	NFO Exchange = "nfo"
	BFO Exchange = "bfo"
	MCX Exchange = "mcx"
)

// Creating methods on enum
func (e Exchange) isEquity() bool {
	return e == NSE || e == BSE
}

func (e Exchange) isDerivatives() bool {
	return e == NFO || e == BFO
}

func (e Exchange) isCommodity() bool {
	return e == MCX
}

type scrip struct {
	token    string
	exchange Exchange
}

func main() {
	//
	scrip := scrip{
		token:    "123",
		exchange: MCX,
	}

	if scrip.exchange.isEquity() {
		fmt.Println("Equity scrip token:", scrip.token)
	} else if scrip.exchange.isDerivatives() {
		fmt.Println("Derivatives scrip")
	} else {
		fmt.Println("Commodity scrip")
	}
}

// Note:
// - Using a custom type gives better type saftey
// - Enum values are still underlying integers,
//   so invalid values are technically possible:
//
//   var s Status = 999
//
// - Because enum-like values are custom types,
//   you can define methods on them.
