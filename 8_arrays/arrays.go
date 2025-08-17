package main

import "fmt"

func main() {
	var nums [4]int // all initialize with 0

	//int -> 0, string -> "", bool -> false

	nums[0] = 1
	fmt.Println(nums[0]) // 1

	fmt.Println(nums)

	fmt.Println(len(nums)) // array length

	var vals [4]bool           //all initialize with false
	vars := [4]int{1, 2, 3, 4} // initialize with values
	vals[2] = true
	fmt.Println(vals)
	fmt.Println(vars)

	//2d arrays
	nums1 := [2][2]int{{1, 2}, {3, 4}}
	//or
	nums2 := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}

	fmt.Println(nums1)
	fmt.Println(nums2)

	// - fixed size, that is predictable
	// - Memory optimization
	// - Oontant time access

}
