package main

import "fmt"

//for ---> only construct in go for looping

func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite
	// for {
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	// we can use "break" and "continue" if needed

	//1.22V    range
	for i := range 3 {
		fmt.Println(i)
	}
}
