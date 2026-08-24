package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for url := range jobs {
		fmt.Printf("Worker %d -> %s\n", id, url)

		resp, err := client.Get(url)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Worker %d -> Status: %d\n", id, resp.StatusCode)

		resp.Body.Close()
	}
}

func main() {
	urls := []string{
		"https://api.qdrant.tech/api-reference/snapshots/create-snapshot",
		"https://huggingface.co/learn/llm-course/chapter6/3b",
		"https://civilica.com/doc/1/",
		"https://derclub.ir/",
		"https://fa.wikipedia.org/wiki/%D8%B5%D9%81%D8%AD%D9%87%D9%94_%D8%A7%D8%B5%D9%84%DB%8C",
	}

	jobs := make(chan string)

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)

		go worker(i, jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	fmt.Println("Done.")
}
