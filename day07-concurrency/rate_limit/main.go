package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://api.qdrant.tech/api-reference/snapshots/create-snapshot",
		"https://derclub.ir/",
		"https://civilica.com/l/1436/",
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for _, url := range urls {
		<-ticker.C

		fmt.Println("Request:", url)

		resp, err := client.Get(url)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Status:", resp.StatusCode)

		resp.Body.Close()
	}
}
