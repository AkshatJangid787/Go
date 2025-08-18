package main

import "fmt"

// iterating over data structure
func main() {
	nums := []int{6, 7, 8}
	sum := 0

	for indx, num := range nums {
		sum += num
		println(indx, num)
	}

	fmt.Println("sum:", sum)

	// map
	m := map[string]string{"frame": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte , 2 byte
	for c, i := range "golang" {
		fmt.Println(c, i)
	}

	for c, i := range "golang" {
		fmt.Println(c, string(i))
	}
}
