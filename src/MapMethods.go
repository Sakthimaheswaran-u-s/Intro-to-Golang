package main

import "fmt"

func main() {
	var sample = make(map[int]string)

	sample[1] = "sakthi"
	sample[2] = "mahes" // intialising the value of the map
	sample[3] = "waran"

	// var a  = sample[3] // accessing value from map
	fmt.Println(sample)

	sample[4] = "us" // updating and adding map elements

	fmt.Println(sample)

	delete(sample, 4) // deleting a particular element from map

	fmt.Println(sample)

	val1, ok1 := sample[4] // checking whether the element with key 2 exits val1 refers to the value of the key 2 and ok1 refers to respective boolean value if exists true else false

	fmt.Println(val1, ok1)

}

// creating map using make function

/*Syntax
var a = make(map[key_type]value_type);
b := make(map[key_type]value_type); */

// deleting the map element using delete function

/* Syntax
delete(map_name, key) */

// checking for specific elements in a map

/*Syntax
val, ok :=map_name[key] */
