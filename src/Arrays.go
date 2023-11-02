package main

import (
	"fmt"
)

var a = [2]string{"test1", "test2"}

func main() {
	var arr = [...]int{1, 2, 3}
	fmt.Println(arr[0], len(a))
}
