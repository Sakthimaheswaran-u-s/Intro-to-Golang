package main

import (
	"fmt"
)

type sample struct {
	name  string
	id    int
	phone int
}

func main() {
	var struct_access sample
	struct_access.name = "sakthi"
	struct_access.id = 1234567
	struct_access.phone = 9876544321
	//  fmt.Println(struct_access);
	JustPrint(struct_access) // passing the reference of the struct as argument to other function
}

func JustPrint(sam sample) { // get the reference along with the structure declared
	fmt.Println(sam.name)
	fmt.Println(sam.id)
	fmt.Println(sam.phone)
}
