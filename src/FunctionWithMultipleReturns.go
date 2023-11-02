package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	a, b := myFunction(5, "Hello") // holding the values in their respective variables
	fmt.Println(a, b)
	// fmt.Println(myFunction(1,"hello"));
}
