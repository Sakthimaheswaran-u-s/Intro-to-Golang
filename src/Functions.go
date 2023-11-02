package main

import (
	"fmt"
)

// func sample()  {
// 	fmt.Println("hello")	// custom function
// }

func sample1(fname string, number int) {
	fmt.Println(fname) // func with parameters and multiple parameters
	fmt.Println(number)
}
func main() {
	sample1("sample print", 2)
}
