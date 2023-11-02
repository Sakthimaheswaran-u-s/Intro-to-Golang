package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(j) // we can add break and continue as usual
		}
	}
}
