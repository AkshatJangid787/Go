package main

import (
	"fmt"
	"slices"
)

// slices are dynamic arrays
// slices are passed by reference
// slices are not passed by value
// slices are not passed by pointer
// + usefull methods

func main() {
	//uninitialized slice is nill
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)

	fmt.Println(len(nums))

	var nums1 = make([]int, 0, 5) //[1 2 3]
	// or
	// nums1  := []int{}

	nums1 = append(nums1, 1)
	nums1 = append(nums1, 2)
	nums1 = append(nums1, 3)

	fmt.Println(nums1)
	fmt.Println(cap(nums1)) //capacity maximum number of fit
	fmt.Println(len(nums1))

	nums1 = append(nums1, 4)

	var nums2 = make([]int, len(nums1))

	//copy function

	copy(nums2, nums1)

	fmt.Println(nums1, nums2)

	// slice operator
	var nums3 = []int{1, 2, 3}
	fmt.Println(nums3[0:2])
	// or
	fmt.Println(nums3[:2])

	fmt.Println(nums3[1:])

	// slice
	var nums4 = []int{1, 2, 3}
	var nums5 = []int{1, 2, 3}

	fmt.Println(slices.Equal(nums4, nums5))

	//2d slices
	var nums6 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums6)
}
