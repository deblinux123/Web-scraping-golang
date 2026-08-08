package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "https://civilica.com/l/1/"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server returned:", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	fmt.Println(string(body))
}
