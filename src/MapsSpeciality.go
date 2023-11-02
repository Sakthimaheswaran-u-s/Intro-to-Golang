package main

import "fmt"

func main() {
	sample := map[string]string{"name": "sak", "initial": "us"}
	fmt.Println(sample)
	sample_1 := sample              // coping the map values from one to other new map
	sample_1["last name"] = "waran" // we change in new map it will also affect the old map, so here the old map is also added with new key value pair (i.e) {"last name" : "waran"}
	fmt.Println(sample_1)
	fmt.Println(sample)
}
