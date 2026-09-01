package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mxschmitt/playwright-go"
)

type FileInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DownloadInfo struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Size   int64  `json:"size"`
}

func main() {
	targetURL := "https://huggingface.co/Qwen/Qwen2.5-0.5B"

	fmt.Println("Target:", targetURL)

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

	fmt.Println("Opening repository ...")

	_, err = page.Goto(
		targetURL,
		playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository opened.")

	filesURL, err := findFilesPage(page)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files page:")
	fmt.Println(filesURL)

	_, err = page.Goto(
		filesURL,
		playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files page opened.")

	files, err := findFiles(page)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFound %d possible files:\n\n", len(files))

	for i, file := range files {
		fmt.Printf("[%d] %s\n", i, file.Name)
		fmt.Printf("		%s\n", file.URL)
	}

	os.Mkdir("downloads", 0755)

	var results []DownloadInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name, "/") {
			continue
		}

		result := downloadFile(file)

		results = append(results, result)
	}

	saveMetadata(results)

	fmt.Println("\nFinished.")
}

func findFilesPage(page playwright.Page) (string, error) {
	links, err := page.Locator("a").All()

	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		href, err := link.GetAttribute("href")

		if err != nil {
			continue
		}

		if href == "" {
			continue
		}

		if strings.Contains(href, "/tree/") {
			baseURL, err := url.Parse(page.URL())

			if err != nil {
				return "", err
			}

			return baseURL.ResolveReference(
				&url.URL{Path: href},
			).String(), nil
		}
	}

	return "", fmt.Errorf("Files page not found")
}

func findFiles(page playwright.Page) ([]FileInfo, error) {
	links, err := page.Locator("a").All()

	if err != nil {
		return nil, err
	}

	var files []FileInfo

	seen := make(map[string]bool)

	for _, link := range links {
		href, err := link.GetAttribute("href")

		if err != nil {
			continue
		}

		if href == "" {
			continue
		}

		if !strings.Contains(href, "/blob/") {
			continue
		}

		text, err := link.InnerText()

		if err != nil {
			continue
		}

		name := strings.TrimSpace(text)

		if name == "" {
			continue
		}

		if seen[href] {
			continue
		}

		seen[href] = true

		baseURL, err := url.Parse(page.URL())

		if err != nil {
			continue
		}

		fullURL := baseURL.ResolveReference(
			&url.URL{Path: href},
		).String()

		files = append(files, FileInfo{
			Name: name,
			URL:  fullURL,
		})
	}

	return files, nil
}

func isLargeModelFile(name string) bool {
	largeExtensions := []string{
		".safetensors",
		".bin",
		".gguf",
		".pt",
		".pth",
	}

	lower := strings.ToLower(name)

	for _, ext := range largeExtensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}

	return false
}

func downloadFile(file FileInfo) DownloadInfo {
	result := DownloadInfo{
		Name:   file.Name,
		URL:    file.URL,
		Status: "failed",
	}

	fmt.Println("\nDownloading:", file.Name)

	downloadURL := strings.Replace(
		file.URL,
		"/blob/",
		"/resolve",
		1,
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(downloadURL)

	if err != nil {
		fmt.Println("Download Error:", err)

		return result
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP status:", resp.Status)

		return result
	}

	outputPath := filepath.Join(
		"download",
		filepath.Base(file.Name),
	)

	output, err := os.Create(outputPath)

	if err != nil {
		fmt.Println("File creation error:", err)

		return result
	}

	defer output.Close()

	size, err := io.Copy(output, resp.Body)

	if err != nil {
		fmt.Println("Copy error:", err)

		return result
	}

	fmt.Printf("Downloaded %s (%d bytes)\n", file.Name, size)

	result.Status = "downloaded"
	result.Size = size

	return result
}

func saveMetadata(results []DownloadInfo) {

	data, err := json.MarshalIndent(
		results,
		"",
		"    ",
	)

	if err != nil {
		fmt.Println("Metadata error:", err)
		return
	}

	err = os.WriteFile(
		"metadata.json",
		data,
		0644,
	)

	if err != nil {
		fmt.Println("Metadata write error:", err)
		return
	}

	fmt.Println("\nMetadata saved to metadata.json")
}
