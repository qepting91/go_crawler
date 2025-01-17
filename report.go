package main

import (
	"fmt"
	"sort"
)

type pageResult struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
  REPORT for %s
=============================
`, baseURL)

	// Convert map to slice for sorting
	var results []pageResult
	for url, count := range pages {
		results = append(results, pageResult{url, count})
	}

	// Sort by count (descending) and URL (ascending)
	sort.Slice(results, func(i, j int) bool {
		if results[i].count != results[j].count {
			return results[i].count > results[j].count
		}
		return results[i].url < results[j].url
	})

	// Print sorted results
	for _, result := range results {
		fmt.Printf("Found %d internal links to %s\n", result.count, result.url)
	}
}
