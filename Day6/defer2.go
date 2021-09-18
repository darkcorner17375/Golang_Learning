package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("Defer.txt")
	defer file.Close() //執行後再關閉
	if err != nil {
		fmt.Println("文件打開錯誤!!", err)
	}
	file.WriteString("我愛Golang!!")
	file.Close()
}
