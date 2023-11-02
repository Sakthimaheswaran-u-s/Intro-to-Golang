package main

import "fmt"

func main() {
	sample_map := map[int]string{1: "name", 2: "initial", 3: "last_name"}

	for key, value := range sample_map { // iteration without any order
		fmt.Println(key, value)
	}
	arr := [...]int{1, 3, 2} // initialising the order

	for _, val := range arr { // iteration with specific order
		fmt.Println(sample_map[val])
	}
}
