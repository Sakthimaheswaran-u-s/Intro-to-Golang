package main

import (
	"fmt"
)

func rec(a int) (b int) {
	if a > 0 {
		b = a * rec(a-1)
	} else {
		b = 1
	}
	return
}
func main() {
	fmt.Println(rec(5))
}
