package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type ProductResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

func main() {
	baseURL := "https://dummyjson.com/products?limit=20&skip=%d"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for page := 1; page <= 5; page++ {
		skip := (page - 1) * 20
		url := fmt.Sprintf(baseURL, skip)

		fmt.Println("page:", page)
		fmt.Println("URL:", url)

		req, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)

		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0",
		)

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Request faile:", err)
			continue
		}

		var data ProductResponse

		err = json.NewDecoder(resp.Body).Decode(&data)

		if err != nil {
			fmt.Println("JSON decode error:", err)
			continue
		}

		fmt.Println("Status:", resp.Status)
		fmt.Println("Products:", len(data.Products))
		resp.Body.Close()

	}
}
