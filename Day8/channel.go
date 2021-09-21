package main

import (
	"fmt"
	"strconv"
)

func main() {
	//宣告頻道
	message := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			message <- (strconv.Itoa(i) + ".hellow channel!")
		}
	}()

	//接收頻道發送的消息
	result := ""
	result = <-message
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
}
