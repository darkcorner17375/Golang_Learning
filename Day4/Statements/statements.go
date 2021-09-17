package main

import "fmt"

func main() {

	//
	if true {
		fmt.Println("This test is true")
	}

	//簡易範例
	score := 30
	if score >= 30 {
		fmt.Println("MVP球員")
	} else if score >= 20 {
		fmt.Println("明星球員")
	} else if score >= 20 {
		fmt.Println("首發球員")
	} else {
		fmt.Println("板凳球員")
	}

	//範例2
	number := 50
	guess := 50

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("Got it!")
		}
	}
}
