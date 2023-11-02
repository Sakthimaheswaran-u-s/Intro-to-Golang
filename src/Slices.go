package main

import (
	"fmt"
)

func main() {
	// sampleSlice := []int{1,2,3};
	// fmt.Println(sampleSlice);

	// creating a slice from an array
	arr1 := [3]int{1, 2, 3}
	sampleSlice := arr1[0:2]
	fmt.Println(sampleSlice)
}

// slices are similar to array but they can grow and shrink their size based on the requirements
// make() can be also used to create slice
// syntax slice_name := make([]type ,length, capacity);
// len() is used to return the length of the slice
// cap() is used to return the capacity of the slice
