package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string){
	baseURL, _ = normalizeURL(baseURL)
	fmt.Println("=============================")
	fmt.Printf("REPORT for https://%s\n", baseURL)
	fmt.Println("=============================")

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return pages[keys[i]] > pages[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("Found %d internal links to %s\n", pages[k], k)
	}
}