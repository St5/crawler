package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages		   int
}

func(cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	pages := make(map[string]int)
	mu := &sync.Mutex{}
	concurrencyControl := make(chan struct{}, maxConcurrency)
	wg := &sync.WaitGroup{}

	return &config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 mu,
		concurrencyControl: concurrencyControl,
		wg:                 wg,
		maxPages: 			maxPages,
	}, nil
}