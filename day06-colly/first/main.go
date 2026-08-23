package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set(
			"User-Agent",
			"Mozilla/5.0",
		)

		fmt.Println("Visiting:", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
		fmt.Println("URL:", r.Request.URL)
		fmt.Println("Status code:", r.StatusCode)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status: ", r.StatusCode)
		fmt.Println("Response size:", len(r.Body))
	})

	c.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println("h1:", h.Text)
	})

	// c.OnHTML("a[href]", func(h *colly.HTMLElement) {
	// 	href := h.Attr("href")

	// 	fmt.Println("Link found:", href)
	// })

	err := c.Visit("https://derclubvv.ir/")

	if err != nil {
		fmt.Println("Error:", err)
	}
}
