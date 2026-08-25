package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
}

func worker(
	id int,
	jobs <-chan string,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for url := range jobs {
		fmt.Printf("Worker %d → %s\n", id, url)

		resp, err := client.Get(url)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		results <- Result{
			URL:    url,
			Status: resp.StatusCode,
		}

		resp.Body.Close()

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	urls := []string{
		"https://derclub.ir/",
		"https://civilica.com/l/1436/",
		"https://qdrant.tech/",
	}

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)

		go worker(i, jobs, results, &wg)
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result.URL, result.URL)
	}

	fmt.Println("Done.")
}
