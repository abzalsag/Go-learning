package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num < 0 {

		fmt.Println("num is negative")
	} else if num == 0 {
		fmt.Println("num is zero")
	} else if num > 0 {
		fmt.Println("num is positive")
	}
}
