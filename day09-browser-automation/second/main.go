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

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			ExecutablePath: playwright.String("/snap/bin/chromium"),
			Headless:       playwright.Bool(false),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	defer browser.Close()

	page, err := browser.NewPage()

	if err != nil {
		log.Fatal(err)
	}

	_, err = page.Goto("https://the-internet.herokuapp.com/dynamic_loading/1")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Page loaded")

	startButton := page.Locator("#start button")

	err = startButton.Click()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start button clicked")

	result := page.Locator("#finish h4")

	err = result.WaitFor(playwright.LocatorWaitForOptions{
		State: playwright.WaitForSelectorStateVisible,
	})

	if err != nil {
		log.Fatal(err)
	}

	txt, err := result.TextContent()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", txt)

}
