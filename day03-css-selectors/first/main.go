package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	html := `
	<div class="product">
		<h2>Laptop</h2>
		<p class="price">1200$</p>
		<a href="/products/1">View</a>
	</div>

	<div class="product">
		<h2>Phone</h2>
		<p class="price">800$</p>
		<a href="/products/2">View</a>
	</div>
	`

	doc, err := goquery.NewDocumentFromReader(
		strings.NewReader(html),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Products:", doc.Find(".product").Length())

	doc.Find(".product h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Product:", s.Text())
	})

	doc.Find(".product .price").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Price:", s.Text())
	})

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists {
			fmt.Println("URL:", href)
		}
	})

	doc.Find(".product > h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
