package main

import "fmt"

/*
func main() {

	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr)

	slice := []string{
		"BMW",
		"Mercedes",
		"Audi",
	}

	slice = append(slice, "Toyota")

	fmt.Println(slice)

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
}

*/

/*
func main() {

	cities := []string{
		"Astana",
		"Moscow",
		"Almaty",
	}

	for _, city := range cities {
		fmt.Println(city)
	}
}

*/

func main() {

	nums := []int{5, 8, 10, 2, 17}

	max := nums[0]

	for _, n := range nums {

		if n > max {
			max = n
		}

	}

	fmt.Println(max)
}
