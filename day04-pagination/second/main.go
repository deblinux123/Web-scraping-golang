package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for id := 1024; id <= 1034; id++ {
		url := fmt.Sprintf("https://civilica.com/l/%d/", id)

		fmt.Println("Request:", url)

		req, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)

		if err != nil {
			fmt.Println("Request creation error:", err)
			continue
		}

		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0",
		)

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Request error:", err)
			continue
		}

		fmt.Println("Status:", resp.Status)

		filename := fmt.Sprintf("page-%d.html", id)

		file, err := os.Create(filename)

		if err != nil {
			fmt.Println("File creation error:", err)
			resp.Body.Close()
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			fmt.Println("Parser error:", err)
			continue
		}

		script := doc.Find("#__INITAL_FROM_SERVER__")

		fmt.Println("Script count:", script.Length())

		doc.Find("div.p-2").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("%d: %s \n", i+1, strings.TrimSpace(s.Text()))
		})

		_, err = file.ReadFrom(resp.Body)

		if err != nil {
			fmt.Println("Save error:", err)
		}

		file.Close()

		resp.Body.Close()

		fmt.Println("Saved:", filename)
	}
}
