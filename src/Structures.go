package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hi");

	type sample struct { // define the struct
		name    string
		roll_no int
		gender  string
	}

	var struct_access sample
	struct_access.name = "hedge" // intialize the values
	struct_access.roll_no = 17
	struct_access.gender = "Male"
	fmt.Println(struct_access)

	/*Syntax
	type struct_name struct{
		member_1 datatype;
		member_2 datatype;
		member_3 datatype;
	}  is used to store values of different data types*/
}
