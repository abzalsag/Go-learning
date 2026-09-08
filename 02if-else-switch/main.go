package main

import "fmt"

func main() {
	/* student := "Abzal"
	age := 19
	if age < 18 {
		fmt.Println(student, "age is less than 18")
	} else {
		fmt.Println(student, "age is greater than 18")
	}
	*/

	/*	var point int
		fmt.Println("point = ")
		fmt.Scanln(&point)
		if point <= 49 {
			fmt.Println("F")
		} else if point <= 69 {
			fmt.Println("D")
		} else if point <= 79 {
			fmt.Println("C")
		} else if point <= 89 {
			fmt.Println("B")
		} else if point <= 99 {
			fmt.Println("A")
		}
	*/
	var day int
	fmt.Println("enter day")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("error day")
	}

}
