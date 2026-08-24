package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		results []string
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			result := fmt.Sprintf("Result %d", id)

			mu.Lock()

			results = append(results, result)

			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(results)
}
