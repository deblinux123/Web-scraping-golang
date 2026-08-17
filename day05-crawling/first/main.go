package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type QueueItem struct {
	URL   string
	Depth int
}

func main() {

	startURL := "https://en.wikipedia.org/wiki/Go_(programming_language)"

	maxDepth := 1

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	baseURL, err := url.Parse(startURL)
	if err != nil {
		panic(err)
	}

	visited := make(map[string]bool)

	queue := []QueueItem{
		{
			URL:   startURL,
			Depth: 0,
		},
	}

	for len(queue) > 0 {

		// Take first item from queue
		item := queue[0]
		queue = queue[1:]

		// Check depth
		if item.Depth > maxDepth {
			continue
		}

		// Check visited
		if visited[item.URL] {
			continue
		}

		visited[item.URL] = true

		fmt.Printf(
			"\n[Crawling] Depth=%d URL=%s\n",
			item.Depth,
			item.URL,
		)

		// HTTP request
		req, err := http.NewRequest(
			http.MethodGet,
			item.URL,
			nil,
		)

		if err != nil {
			fmt.Println("Request error:", err)
			continue
		}

		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/151.0 Safari/537.36",
		)

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Request error: ", err)
			continue
		}

		// Always close body
		defer resp.Body.Close()

		// Check status
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Status:", resp.StatusCode)
			continue
		}

		// Check Content-Type
		contentType := resp.Header.Get("Content-Type")

		if !strings.Contains(contentType, "text/html") {
			fmt.Println("Not HTML:", contentType)
			continue
		}

		// Parse HTML
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		// Extract links
		doc.Find("a").Each(func(i int, s *goquery.Selection) {

			href, exists := s.Attr("href")

			if !exists || href == "" {
				return
			}

			// Parse link
			linkURL, err := url.Parse(href)
			if err != nil {
				return
			}

			// Convert relative URL to absolute URL
			absoluteURL := baseURL.ResolveReference(linkURL)

			// Remove fragment
			absoluteURL.Fragment = ""

			// Domain restriction
			if absoluteURL.Hostname() != baseURL.Hostname() {
				return
			}

			nextURL := absoluteURL.String()

			// Check visited
			if visited[nextURL] {
				return
			}

			fmt.Printf(
				"  [Found] %s\n",
				nextURL,
			)

			// Add to queue
			queue = append(queue, QueueItem{
				URL:   nextURL,
				Depth: item.Depth + 1,
			})
		})
	}

	fmt.Println("\nCrawling finished.")
	fmt.Println("Pages visited:", len(visited))
}
