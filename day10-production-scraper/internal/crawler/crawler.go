package crawler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

type Crawler struct {
	client     *http.Client
	maxRetries int
	delay      time.Duration
	logger     *slog.Logger
}

func New(timeout time.Duration, maxRetries int, delay time.Duration, logger *slog.Logger) *Crawler {
	return &Crawler{
		client: &http.Client{
			Timeout: timeout,
		},
		maxRetries: maxRetries,
		delay:      delay,
		logger:     logger,
	}
}

func (c *Crawler) Fetch(ctx context.Context, url string) ([]byte, error) {
	var lastErr error

	for attempt := 1; attempt <= c.maxRetries; attempt++ {
		c.logger.Info(
			"Fetching URL",
			"URL", url,
			"Attempt", attempt,
		)

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			url,
			nil,
		)

		if err != nil {
			return nil, err
		}

		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0 (compatible; ProductionScraper/1.0)",
		)

		resp, err := c.client.Do(req)

		if err != nil {
			lastErr = err
			c.logger.Warn(
				"Request failed",
				"Error", err,
				"Attemt", attempt,
			)
			c.sleep(ctx, attempt)
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()

		if readErr != nil {
			lastErr = readErr
			c.sleep(ctx, attempt)
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			c.logger.Info(
				"Request successful",
				"URL", url,
				"Status", resp.StatusCode,
			)

			return body, nil
		}

		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusRequestTimeout ||
			resp.StatusCode >= 500 {
			lastErr = fmt.Errorf(
				"Server returned status %d",
				resp.StatusCode,
			)

			c.logger.Warn(
				"Retryable HTTP status",
				"Status", resp.StatusCode,
				"Attemtp", attempt,
			)

			c.sleep(ctx, attempt)
			continue
		}

		return nil, fmt.Errorf(
			"non-retryable HTTP status: %d",
			resp.StatusCode,
		)
	}

	return nil, fmt.Errorf(
		"Request failed after %d attemtps: %w",
		c.maxRetries,
		lastErr,
	)
}

func (c *Crawler) sleep(ctx context.Context, attempt int) {
	backoff := c.delay * time.Duration(1<<(attempt-1))

	jitter := time.Duration(rand.Int63n(int64(c.delay)))

	wait := backoff + jitter

	c.logger.Info(
		"Waiting befor retry",
		"Duration", wait,
	)

	timer := time.NewTimer(wait)

	defer timer.Stop()

	select {
	case <-ctx.Done():
	case <-timer.C:
	}
}
