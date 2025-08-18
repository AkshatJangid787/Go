package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	//creating map

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	//get an element
	fmt.Println(m["name"], m["area"])

	//if key does not exists in the map then it returns zero
	fmt.Println(m["phone"])

	m1 := make(map[string]int)
	m1["age"] = 30
	m1["dob"] = 14
	fmt.Println(m1["phone"]) // it will return zero

	//int->0, string->empty, bool->false

	fmt.Println(len(m))

	delete(m, "area") // delete particular key
	clear(m)          //map will empty
	fmt.Println(m)

	// or

	m2 := map[string]int{"price": 40, "phone": 3}
	fmt.Println(m2)

	v, ok := m2["price"]
	fmt.Println(v) // it will return value of that key --> 40

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(m1, m2)) // compare two maps (true/false)

}
