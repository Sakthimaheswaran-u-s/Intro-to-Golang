package main

import (
	"fmt"
)

var a string = "32.90"

const (
	value  = 32
	value1 = 34
)

func main() {
	// var a bool = true ;
	var va = (value1)
	fmt.Printf("%T", va)
}

// %v is value
// %T is find data type
// %% is used to print the value with percentage sign
// %b is used to print number with base 2
// %4d	Pad with spaces (width 4, right justified)
// %-4d	Pad with spaces (width 4, left justified)
// %04d	Pad with zeroes (width 4)
// %q	Prints the value as a double-quoted string
// %8s	Prints the value as plain string (width 8, right justified)
// %-8s	Prints the value as plain string (width 8, left justified)
