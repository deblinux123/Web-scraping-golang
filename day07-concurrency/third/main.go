package main

import (
	"fmt"
	"sync"
	"time"
)

func even(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	if number%2 == 0 {
		fmt.Println("Number is even:", number)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Finished processing.")
}

func main() {
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5, 6, 12, 445, 55, 32, 43}

	for _, num := range numbers {
		wg.Add(1)

		go even(num, &wg)
	}

	wg.Wait()

	fmt.Println("All numbers are processed.")
}
