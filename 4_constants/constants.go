package main

import "fmt"

const pie = 3.14

//name := "akshat" ---> shorthand syntax can't declare global

func main() {
	const name string = "golang"
	const age = 21

	fmt.Println(pie)
	fmt.Println(name)
	fmt.Println(age)

	const (
		port = 5000
		host = "LocalHost"
	)

	fmt.Println(port, host)
}
