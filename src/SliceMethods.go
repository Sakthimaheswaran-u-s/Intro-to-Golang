package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 2, 3, 4, 5}

	/* Accessing the element from slice
	fmt.Println(prices[4]);*/

	/* Changing the value of the element in slice
	prices[2] = 40; */

	/* adding elements to end of the slices
	   prices = append(prices, 20, 40);*/

	/*  appending one slice to other slices
	sampleSlice := []int{1,2,3,4,5};
	pricesAppend := append(prices, sampleSlice...);
	fmt.Println(pricesAppend);
	The '...' after slice2 is necessary when appending the elements of one slice to another.*/

	fmt.Println(prices)

}
