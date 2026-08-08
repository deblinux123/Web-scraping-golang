# 🕷️ Go Web Scraping — 10 Day Learning Journey

A practical **10-day journey to learn Web Scraping and Web Crawling with Go (Golang)** — starting from HTTP fundamentals and HTML parsing, and gradually moving toward concurrency, dynamic websites, browser automation, and production-ready scrapers.

The goal of this repository is not just to learn scraping libraries, but to understand **how web scraping works under the hood** and how to build reliable, scalable scrapers with Go.

---

## 🎯 Goals

By the end of this 10-day journey, you should be able to:

* Understand HTTP requests and responses
* Work with Go's `net/http` package
* Parse HTML documents
* Use CSS selectors
* Extract structured data from websites
* Handle pagination
* Build web crawlers
* Use Colly for advanced scraping
* Build concurrent scrapers with Goroutines and Channels
* Work with websites that expose APIs
* Understand dynamic JavaScript websites
* Automate browsers when necessary
* Store scraped data in databases
* Handle retries, timeouts, rate limits, and errors
* Build a production-oriented web scraper

---

# 🗓️ 10-Day Roadmap

| Day    | Topic                   | Main Concepts                              |
| ------ | ----------------------- | ------------------------------------------ |
| **01** | HTTP & `net/http`       | Requests, Responses, Headers, Status Codes |
| **02** | HTML Parsing            | DOM, HTML, `goquery`                       |
| **03** | CSS Selectors           | Selectors, Attributes, Nested Elements     |
| **04** | Pagination              | Multiple Pages, Next Page, Data Collection |
| **05** | Web Crawling            | Links, Depth, Visited URLs, Domains        |
| **06** | Colly                   | Collector, Callbacks, Requests, Responses  |
| **07** | Concurrency             | Goroutines, Channels, Worker Pools         |
| **08** | APIs & Dynamic Websites | REST APIs, JSON, Network Requests          |
| **09** | Browser Automation      | JavaScript, Playwright, Chromium           |
| **10** | Production Scraper      | Database, Logging, Retry, Rate Limiting    |

---

# 📂 Project Structure

```text
go-web-scraping-10-days/
│
├── README.md
├── go.mod
├── .gitignore
│
├── day01-http-basics/
│   ├── main.go
│   └── README.md
│
├── day02-html-parsing/
│   ├── main.go
│   └── README.md
│
├── day03-css-selectors/
│   ├── main.go
│   └── README.md
│
├── day04-pagination/
│   ├── main.go
│   └── README.md
│
├── day05-crawling/
│   ├── main.go
│   └── README.md
│
├── day06-colly/
│   ├── main.go
│   └── README.md
│
├── day07-concurrency/
│   ├── main.go
│   └── README.md
│
├── day08-apis-dynamic-websites/
│   ├── main.go
│   └── README.md
│
├── day09-browser-automation/
│   ├── main.go
│   └── README.md
│
└── day10-production-scraper/
    ├── cmd/
    │   └── scraper/
    │       └── main.go
    │
    ├── internal/
    │   ├── scraper/
    │   ├── parser/
    │   └── storage/
    │
    └── README.md
```

---

# 🟢 Day 01 — HTTP & `net/http`

Learn how web communication works.

### Topics

* HTTP Request
* HTTP Response
* HTTP Methods
* Status Codes
* Headers
* `http.Client`
* Request timeout
* Response body
* Error handling

### First scraper

```text
URL
 ↓
HTTP Request
 ↓
HTTP Response
 ↓
HTML
 ↓
Save / Print HTML
```

Main package:

```go
net/http
```

---

# 🟢 Day 02 — HTML Parsing

Learn how to extract information from HTML.

### Topics

* HTML
* DOM
* Elements
* Attributes
* Text
* `goquery`

Example:

```html
<div class="product">
    <h2>Laptop</h2>
    <span class="price">$1000</span>
</div>
```

Extract:

```text
Laptop
$1000
```

---

# 🟢 Day 03 — CSS Selectors

Learn how to find specific elements inside the DOM.

Examples:

```css
.product
.product h2
.product .price
a
a[href]
div.product > h2
```

You will learn how to combine selectors to extract structured data reliably.

---

# 🟡 Day 04 — Pagination

Learn how to scrape multiple pages.

Example:

```text
/products?page=1
/products?page=2
/products?page=3
...
/products?page=100
```

Goal:

```text
100 pages
×
20 products
=
2000 products
```

Save the results as:

```text
JSON
CSV
```

---

# 🟡 Day 05 — Web Crawling

Move from scraping a single page to crawling an entire website.

Learn:

* Extracting links
* Following links
* Visited URLs
* Duplicate detection
* Domain restriction
* Crawl depth
* Crawl queue

Architecture:

```text
Homepage
   │
   ├── Page A
   │      ├── Page C
   │      └── Page D
   │
   └── Page B
          └── Page E
```

---

# 🟠 Day 06 — Colly

Learn how to use **Colly**, a powerful web scraping and crawling framework for Go.

Topics:

```text
Collector
OnHTML
OnRequest
OnResponse
OnError
Visit
```

Build a real scraper using Colly.

---

# 🔴 Day 07 — Concurrency

Use Go's concurrency model to make scraping faster.

Learn:

* Goroutines
* Channels
* `sync.WaitGroup`
* Mutex
* Worker Pools
* Concurrent requests
* Rate limiting

Architecture:

```text
                    ┌── Worker 1
                    │
URL Queue ──────────┼── Worker 2
                    │
                    ├── Worker 3
                    │
                    └── Worker 4
```

---

# 🔴 Day 08 — APIs & Dynamic Websites

Not every website should be scraped from HTML.

Learn how to inspect:

```text
Browser
   ↓
DevTools
   ↓
Network
   ↓
Fetch / XHR
   ↓
API
   ↓
JSON
```

Topics:

* REST APIs
* JSON
* HTTP API requests
* Network inspection
* API endpoints
* Request headers
* Query parameters

Whenever a website exposes the data through an accessible API, prefer using the API instead of unnecessarily automating a browser.

---

# 🔴 Day 09 — Browser Automation

Some websites require JavaScript execution.

Learn browser automation with tools such as:

* Playwright
* chromedp
* Chromium

Architecture:

```text
Go
 │
 ▼
Browser Automation
 │
 ▼
Chromium
 │
 ├── JavaScript
 ├── Cookies
 ├── DOM
 └── Network
```

The goal is to understand **when browser automation is actually necessary**.

---

# 🚀 Day 10 — Production Scraper

Combine everything into a production-oriented scraper.

Architecture:

```text
                    ┌──────────────┐
                    │   Website    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   Crawler    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │    Parser    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │    Cleaner   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   Storage    │
                    └──────────────┘
```

Features:

* Concurrent workers
* Rate limiting
* Retry mechanism
* Timeout handling
* Logging
* Error handling
* Duplicate detection
* Data validation
* PostgreSQL storage
* Configuration management

---

# 🧰 Technologies

The project will progressively use:

```text
Go
 │
 ├── net/http
 ├── goquery
 ├── Colly
 ├── Goroutines
 ├── Channels
 ├── Playwright / chromedp
 ├── PostgreSQL
 └── Redis
```

---

# 📚 Learning Philosophy

This repository follows one important rule:

> **Understand the fundamentals before using frameworks.**

For example, before using Colly, we first build scrapers using:

```text
net/http
      +
goquery
```

This helps us understand what scraping frameworks are actually doing for us.

The same principle applies to concurrency and browser automation.

---

# 🧪 Projects

During the 10 days, the exercises will gradually become more complex.

### Project 1

Simple HTML scraper:

```text
Title
Price
URL
```

### Project 2

Pagination scraper:

```text
100+ pages
```

### Project 3

Website crawler:

```text
Homepage
 ↓
Links
 ↓
Pages
 ↓
More links
```

### Project 4

Concurrent scraper:

```text
URL Queue
 ↓
Worker Pool
 ↓
Scrapers
 ↓
Results
```

### Final Project

Production-oriented data scraper:

```text
Website
   ↓
Crawler
   ↓
Concurrent Workers
   ↓
Parser
   ↓
Cleaner
   ↓
PostgreSQL
   ↓
API
```

---

# ⚠️ Responsible Scraping

Web scraping should be performed responsibly.

Before scraping a website:

* Check its Terms of Service.
* Respect `robots.txt` where applicable.
* Avoid excessive request rates.
* Respect rate limits.
* Avoid collecting sensitive personal information.
* Do not attempt to bypass authentication or access controls.
* Prefer official APIs when available.
* Cache data when possible to reduce unnecessary requests.

The purpose of this project is to learn **ethical and responsible web data extraction**.

---

# 📈 Progress

Track the learning progress here:

* [ ] Day 01 — HTTP & `net/http`
* [ ] Day 02 — HTML Parsing
* [ ] Day 03 — CSS Selectors
* [ ] Day 04 — Pagination
* [ ] Day 05 — Web Crawling
* [ ] Day 06 — Colly
* [ ] Day 07 — Concurrency
* [ ] Day 08 — APIs & Dynamic Websites
* [ ] Day 09 — Browser Automation
* [ ] Day 10 — Production Scraper

---

# 🎯 Final Goal

By the end of these 10 days, the goal is to be able to look at a website and answer:

> **Where is the data coming from, how can I access it, how should I extract it, and how can I build a reliable Go scraper to collect it at scale?**

This repository is a **hands-on learning journey**, not just a collection of code examples.

---

## 👨‍💻 Author

**Farid**

Learning:

```text
Golang
Backend Development
Web Scraping
Web Crawling
AI / ML
```

---

## ⭐ Progress

If this repository helps you learn Go Web Scraping, feel free to ⭐ the repository and follow the journey.
