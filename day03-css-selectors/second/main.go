package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://civilica.com/l/2/"

	// creating client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// new request
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("request failed:", err)
		return
	}

	// close when its done
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	// get the doc
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("failed to pars html:", err)
		return
	}

	aside := doc.Find("aside")

	fmt.Println("Aside count:", aside.Length())
	fmt.Println("H3 count:", aside.Find("h3").Length())

	name := aside.Find("h3").First().Text()
	fmt.Println("Name:", name)

	associationType := aside.Find("b").First().Text()

	fmt.Println("Type:", associationType)

	fmt.Println("Image count:", aside.Find("img").Length())

	image, exists := aside.Find("img").First().Attr("src")

	if exists {
		fmt.Println("image:", image)
	}

	alt, exists := aside.Find("img").First().Attr("alt")

	if exists {
		fmt.Println("Alt:", alt)
	}

	images := doc.Find("aside img[src]")

	fmt.Println("Image with src:", images.Length())

	src, exists := images.First().Attr("src")

	if exists {
		fmt.Println("Image src:", src)
	}

	links := doc.Find("aside a[href]")

	fmt.Println("A:", links.Length())

	href, exists := links.First().Attr("href")

	if exists {
		fmt.Println("href:", href)
	}

	fmt.Println("1:", doc.Find("aside h3").Length())
	fmt.Println("2:", doc.Find("aside > h3").Length())
	fmt.Println("3:", doc.Find("aside div > h3").Length())

	doc.Find("aside .shadow-me h3").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%d: %s\n", i+1, strings.TrimSpace(s.Text()))
	})
}
