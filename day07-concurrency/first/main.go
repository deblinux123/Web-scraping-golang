package main

import (
	"fmt"
	"time"
)

func scrape(url string) {
	fmt.Println("Scraping:", url)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished:", url)
}

func main() {
	go scrape("https://civilica.com/")
	go scrape("https://huggingface.co/learn/llm-course/chapter6/3b")
	go scrape("https://api.qdrant.tech/api-reference/snapshots/create-snapshot")

	time.Sleep(3 * time.Second)
}
