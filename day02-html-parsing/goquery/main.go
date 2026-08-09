package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Contact struct {
	Address string
	Phone   string
	Fax     string
	Email   string
	Website string
}

func main() {
	url := "https://civilica.com/l/2/"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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

	err = os.WriteFile("civilica.html", body, 0644)

	if err != nil {
		fmt.Println("Failed to save HTML:", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))

	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	title := doc.Find("title").Text()
	fmt.Println("Title:", title)

	selector := "aside"

	aside := doc.Find(selector)

	name := aside.Find("h3").First().Text()

	fmt.Println("Name:", name)

	// conect := aside.Find(".mb-2").Text()
	// fmt.Println("Contact:")
	// fmt.Println(conect)

	associationType := aside.Find("b").First().Text()

	fmt.Println("Type:", associationType)

	contact := aside.Find("div.mt-4")
	var contactInfo Contact

	contact.Contents().Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		if text == "" {
			return
		}

		parts := strings.SplitN(text, ":", 2)

		if len(parts) != 2 {
			return
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {

		case "آدرس":
			contactInfo.Address = value

		case "تلفن":
			contactInfo.Phone = value

		case "فکس":
			contactInfo.Fax = value

		case "ایمیل":
			contactInfo.Email = value

			// case "وبسایت":
			// 	contactInfo.Website = value
		}
	})

	fmt.Println("Address:", contactInfo.Address)
	fmt.Println("Phone:", contactInfo.Phone)
	fmt.Println("Fax:", contactInfo.Fax)
	fmt.Println("Email:", contactInfo.Email)
}
