package main

import (
	"fmt"
)

func main() {
	a := 12
	b := 13
	if a < b {
		fmt.Println(a == b)
	} else if b > a {
		fmt.Println(true)
	} else {
		fmt.Println("Hello")
	}
	fmt.Println("if-else")

	/* Nested If
	if(a){
		if(a>b){
			fmt.Println("Hello");
		}
	}*/
}
