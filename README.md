# Web Crawler

A fast, concurrent web crawler written in Go that maps the structure of websites by following internal links.

## Features

- Concurrent crawling with configurable worker limits
- Page visit tracking and reporting
- Configurable maximum page limits
- Domain-scoped crawling
- Sorted results by link frequency

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go_crawler.git
cd go_crawler# go_crawler
```

2. Initalize the go module:
```bash
go mod init github.com/yourusername/go_crawler
go mod tidy
```
## Usage
1. Run the crawler:
```bash
go run . <url> <maxConcurrency> <maxPages>
```

## Example
```bash
go run . https://example.com 5 100
```
### Parameters:

url: The starting URL to crawl
maxConcurrency: Maximum number of concurrent requests (recommended: 3-10)
maxPages: Maximum number of pages to crawl

## Example Output
```bash
=============================
  REPORT for https://example.com
=============================
Found 5 internal links to /about
Found 3 internal links to /products
Found 2 internal links to /contact
```
