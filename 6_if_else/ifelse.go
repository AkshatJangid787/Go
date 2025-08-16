package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("PERSON IS AN ADULT")
	// } else {
	// 	fmt.Println("NOT AN ADULT")
	// }

	// if age >= 18 {
	// 	fmt.Println("PERSON IS AN ADULT")
	// } else if age >= 12 {
	// 	fmt.Println("TEENAGER")
	// } else {
	// 	fmt.Println("kid")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" || hasPermissions {
	// 	fmt.Println("yes")
	// }

	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Teenager", age)
	} else {
		fmt.Println("kid", age)
	}

	// go does not have ternary, we have to use normal if-else
}
