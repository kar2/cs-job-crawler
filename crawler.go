package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func crawl(numPages int) map[int]string {

	fmt.Println("Starting crawler...")
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// Map job ID to link
	linkMap := make(map[int]string)

	// Check for link on page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		rawLink := e.Attr("href")

		// Filter links for clean job links
		if strings.Contains(rawLink, "/jobs/view/") && strings.Contains(rawLink, "?") && !strings.Contains(rawLink, "externalApply") {

			// Parse link, add to map
			jobLink := rawLink[:strings.Index(rawLink, "?")]
			jobSlice := strings.Split(jobLink, "-")
			jobID, _ := strconv.Atoi(jobSlice[len(jobSlice)-1])
			if linkMap[jobID] == "" && len(linkMap) < numPages {
				linkMap[jobID] = jobLink
				fmt.Println("Visiting: " + jobLink)
				c.Visit(e.Request.AbsoluteURL(jobLink))
			}
		} else {
			return
		}

	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Base url
	c.Visit("https://www.linkedin.com/jobs/view/1404185031/")

	fmt.Println("Shutting down crawler...")
	fmt.Println()

	return linkMap
}
