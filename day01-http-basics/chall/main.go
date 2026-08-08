package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	url := os.Args[1]
	fmt.Println("URl: ", url)
	fmt.Println("Sending request ...")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (compatible; GoWebScraper/1.0)",
	)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Content-Length:", resp.ContentLength)
	fmt.Println("Response Headers:")
	fmt.Println(resp.Header)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server returned a none 200 status code")
		return
	}

	html, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read response boyd:", err)
		return
	}

	err = os.WriteFile(
		"page.html",
		html,
		0644,
	)

	if err != nil {
		fmt.Println("Failed to save HTML:", err)
		return
	}

	fmt.Println("Saved success fully: page.html")
}
