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

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		ExecutablePath: playwright.String("/snap/bin/chromium"),
		Headless:       playwright.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	_, err = page.Goto("https://huggingface.co/Qwen/Qwen2.5-0.5B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Website opened!")

	links, err := page.Locator("a").All()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d links\n\n", len(links))

	total := 0

	for i, link := range links {
		text, _ := link.InnerText()
		href, _ := link.GetAttribute("href")

		total += 1
		fmt.Printf("[%d] TEXT: %q | URL: %q\n", i+1, text, href)
	}

	fmt.Println("Total links:", total)
}
