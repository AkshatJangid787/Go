package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
		// break ----> no need to use "Break" like other languages because "Go" handles it
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// mutiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("i is an int", t)
		case string:
			fmt.Println("i is a string", t)
		case bool:
			fmt.Println("i is a bool", t)
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("goolang")
	whoAmI(1)
	whoAmI(true)

}
