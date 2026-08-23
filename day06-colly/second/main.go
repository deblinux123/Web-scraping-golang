package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// set up a callback for when a request is made
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set(
			"User-Agent",
			"Mozilla/5.0",
		)

		fmt.Println("Visiting:", r.URL.String())
	})

	// get the response and print the status code and response size
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
		fmt.Println("Response size:", len(r.Body))
	})

	// get the text from the id and print the text
	c.OnHTML(".prose-doc p", func(h *colly.HTMLElement) {
		fmt.Println("Text:", h.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	err := c.Visit("https://huggingface.co/learn/llm-course/chapter6/1")

	if err != nil {
		fmt.Println("Error:", err)
	}
}
