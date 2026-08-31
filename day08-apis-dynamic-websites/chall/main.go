package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func fetchProducts(client *http.Client, page int, limit int) ([]Product, error) {

	u, err := url.Parse("https://dummyjson.com/products")

	if err != nil {
		return nil, err
	}

	query := u.Query()

	query.Set("limit", fmt.Sprintf("%d", limit))
	query.Set("skiip", fmt.Sprintf("%d", (page-1)*limit))

	u.RawQuery = query.Encode()

	fmt.Println("Request:", u.String())

	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	req.Header.Set(
		"Accept",
		"application/json",
	)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Unexpected status code: %d", resp.StatusCode,
		)
	}

	var data ProductResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Products, nil
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	const (
		totalPages  = 5
		productPage = 10
	)

	var allProducts []Product

	for page := 1; page <= totalPages; page++ {
		fmt.Println()
		fmt.Println(strings.Repeat("=", 30))
		fmt.Println("Page: ", page)
		fmt.Println(strings.Repeat("=", 30))

		products, err := fetchProducts(client, page, productPage)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for _, product := range products {
			fmt.Printf("ID: %d | Title: %s | Price: %.2f\n", product.ID, product.Title, product.Price)
		}

		allProducts = append(allProducts, products...)
	}

	fmt.Println()
	fmt.Println("========================")
	fmt.Println("SCRAPING FINISHED")
	fmt.Println("========================")

	fmt.Println(
		"Total products:",
		len(allProducts),
	)
}
