package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)



func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("no max concurrency provided")
		os.Exit(1)
	}
	if len(os.Args) < 4 {
		fmt.Println("no max page provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	maxConcurrency,err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("error parsing max concurrency:", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("error parsing max pages:", err)	
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", baseUrl)

	cfg, err := configure(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("error configuring:", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	printReport(cfg.pages, baseUrl)

	elapsed := time.Since(start)
    fmt.Printf("The script took %s \n", elapsed)

}



