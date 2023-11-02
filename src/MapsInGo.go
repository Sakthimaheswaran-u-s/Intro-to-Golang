package main

import "fmt"

func main() {
	// var a = map[string]string{"name":"sakthi","gender":"male", "ID":"I1234"}
	b := map[int]string{1: "sakthi", 2: "ram", 3: "mahes"}
	fmt.Println(b)
}

/*Syntax
var map_name = map[Key_Type]ValueType{key1:value1, key2:value2,...} */

// map is unordered and changeable collection that does not allow duplicate values
// len() is used to find the length of the map
