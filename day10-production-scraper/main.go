package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/deblinux123/Web-scraping-golang/day10-production-scraper/internal/crawler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	c := crawler.New(
		10*time.Second,
		3,
		1*time.Second,
		logger,
	)

	ctx := context.Background()

	body, err := c.Fetch(
		ctx,
		"https://scrapifydatalabs.com/playground/static-products",
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("products.html", body, 0644); err != nil {
		log.Fatal(err)
	}

	logger.Info("page saved", "bytes", len(body))
}
