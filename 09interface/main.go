package main

import "fmt"

/*
type Payment interface {
	Pay()
}

type Kaspi struct{}

func (Kaspi) Pay() {
	fmt.Println("Kaspi Payment")
}

func main() {
	var p Payment
	p = Kaspi{}
	p.Pay()
}

*/

type Printer interface {
	Print()
}

type Document struct{}

func (Document) Print() {

	fmt.Println("Printing...")

}

func main() {

	var p Printer = Document{}

	p.Print()

}
