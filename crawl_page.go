package main

import (
	"fmt"
	"net/url"
)


func (cfg *config) crawlPage(rawCurrentURL string) {
	
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.maxPages <= len(cfg.pages) {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		
		return
	}

	urlNorm, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	isFirst := cfg.addPageVisit(urlNorm)
	if !isFirst {
		return
	}
	

	body, err := getHTML(rawCurrentURL)

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(body)
	fmt.Println(rawCurrentURL)
	//fmt.Println("-------------------")

	listUrls, err := getURLsFromHTML(body, rawCurrentURL)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range listUrls {
		
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}

}