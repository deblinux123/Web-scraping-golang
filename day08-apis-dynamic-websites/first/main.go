package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
}

type ProductResponse struct {
	Products []Product `json:"products"`
}

func main() {
	url := "https://dummyjson.com/products"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("Response error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	var data ProductResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		fmt.Println("Json error:", err)
		return
	}

	for _, product := range data.Products {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("Id:", product.ID)
		fmt.Println("Title:", product.Title)
		fmt.Println("Description:", product.Description)
		fmt.Println("Category:", product.Category)
		fmt.Println("Price:", product.Price)
		fmt.Println("Brand:", product.Brand)
	}
}
