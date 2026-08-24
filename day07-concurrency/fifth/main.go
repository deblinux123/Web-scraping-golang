package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d scraping %s\n", id, url)

		time.Sleep(1 * time.Second)

		fmt.Printf("Worker %d finished %s\n", id, url)
	}
}

func main() {
	urls := []string{
		"https://api.qdrant.tech/api-reference/snapshots/create-snapshot",
		"https://huggingface.co/learn/llm-course/chapter6/3b",
		"https://civilica.com/doc/1/",
	}

	jobs := make(chan string)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go worker(i, jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs finished.")
}
