package main

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw, err := playwright.Run()

	if err != nil {
		log.Fatal(err)
	}

	defer pw.Stop()

	// lunch chromium
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		ExecutablePath: playwright.String("/snap/bin/chromium"),
		Headless:       playwright.Bool(false),
	})

	if err != nil {
		log.Fatal(err)
	}

	defer browser.Close()

	// create new page
	page, err := browser.NewPage()

	if err != nil {
		log.Fatal(err)
	}

	// open website
	_, err = page.Goto("https://civilica.com/l/1436/")

	if err != nil {
		log.Fatal(err)
	}

	html, err := page.Content()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)

	// get page title
	title, err := page.Title()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title", title)

	// get current url
	fmt.Println("URL:", page.URL())

}
