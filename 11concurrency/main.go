package main

import "fmt"

/*
import (
	"fmt"
	"time"
)

func work() {
	fmt.Println("Working...")
}

func main() {
	go work()
	time.Sleep(time.Second)

}

*/

/*
import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printNumbers(&wg)
	wg.Wait()
}
*/

func main() {

	ch := make(chan string)

	go func() {

		ch <- "Hello Go"

	}()

	fmt.Println(<-ch)

}
