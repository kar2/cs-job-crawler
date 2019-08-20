package main

import (
	"flag"
	"fmt"
)

func main() {

	// Parse cmd line arg
	numPagesPtr := flag.Int("n", 25, "Number of jobs to crawl. (Optional)")
	baseURLPtr := flag.String("b", "", "Base URL to start crawling. (Required)")
	flag.Parse()

	// Crawl and process jobs
	if *baseURLPtr == "" {
		fmt.Println("Error: Base URL is required and must be a specific Linkedin job page.")
		return
	}
	linkMap := crawl(*numPagesPtr, *baseURLPtr)
	exportAsTSV(process(linkMap))
}
