package main

/*
import "fmt"


func changeAge(age *int) {
	*age = 20
}

func main() {
	age := 19

	changeAge(&age)

	fmt.Println(age)
}

*/

/*
import "fmt"
import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)

	result2, err := Divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result2)
}

*/

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {

	if age < 18 {
		return errors.New("too young")
	}

	return nil

}

func main() {

	err := checkAge(16)

	fmt.Println(err)

}
