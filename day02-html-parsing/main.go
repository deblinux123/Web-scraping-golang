package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "https://civilica.com/l/31/"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read body:", err)
		return
	}

	fmt.Println("HTML size:", len(body))

}
