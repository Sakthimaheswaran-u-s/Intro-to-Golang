package main

import (
	"fmt"
)

func main() {
	i := 5
	for i > 3 { // We can use for like this instead of while loop
		fmt.Println(i)
		i--
	}
}
