package main

import (
	"fmt"
)

func sampleRecursion(a int) int {
	if a == 3 {
		return 0
	}
	fmt.Println(a)
	return sampleRecursion(a + 1)
}

func main() {
	sampleRecursion(1)
}
