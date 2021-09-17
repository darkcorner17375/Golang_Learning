package main

import (
	"fmt"
	"time"
)

func main() {
	//宣告Switch
	switch i := 2 + 3; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three,four,five")
	default:
		fmt.Println("another number")
	}

	//範例2
	x := 10
	switch {
	case x <= 10:
		fmt.Println("less than or equal to ten")
	case x <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	//範例3
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("周末來了")
	default:
		fmt.Println("NO~~")
	}

}
