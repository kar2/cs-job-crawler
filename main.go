package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Starting...")
	// scrape()

	//baseURL := "https://www.linkedin.com"
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Get new link

		if e.Attr("data-control-name") != "jobdetails_similarjobs" {
			fmt.Println("Ending...")
			return
		}
		link := e.Attr("href")
		fmt.Println("Link to next page: " + link)
		e.Request.Visit(link)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("https://www.linkedin.com/jobs/view/1404185031/")

	fmt.Println("Ending...")
}
