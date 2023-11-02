package main

import (
	"fmt"
)

func add(a int, b int) (res int) { // named return values
	res = a + b
	return
}

func main() {
	sum_value := add(23, 45) // storing return value in a variable
	fmt.Println(sum_value)
}
