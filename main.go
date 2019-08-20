package main

import "flag"

func main() {

	// Parse cmd line arg
	numPagesPtr := flag.Int("n", 25, "Number of jobs to crawl.")
	flag.Parse()

	// Crawl and process jobs
	linkMap := crawl(*numPagesPtr)
	exportAsTSV(process(linkMap))
}
