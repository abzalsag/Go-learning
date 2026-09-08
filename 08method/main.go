package main

import "fmt"

/*  1
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
} */

/* 2 type BankAccount struct {
	Name    string
	Age     int
	Balance int
}

func (b BankAccount) Deposite(amount int) {
	b.Balance += amount

}
*/

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Height + r.Width)
}

func main() {
	// 1  r := Rectangle{23, 10}
	// 1  fmt.Println(r.Area())

	// 2  acc := BankAccount{"Abzal", 19, 1000}
	// 2  fmt.Println(acc)

	r := Rectangle{10, 20}
	fmt.Println(r.Perimeter())
}
