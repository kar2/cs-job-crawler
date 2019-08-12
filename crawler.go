package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func scrape() {
	c := colly.NewCollector(
		// Only allow linkedin.com links
		colly.AllowedDomains("linkedin.com"),

		// Allow async
		colly.Async(true),

		// Max depth 3 for now
		colly.MaxDepth(3),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Get new link
		link := e.Attr("href")

		fmt.Println(e.Text)

		c.Visit(e.Request.AbsoluteURL(link))

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("http://www.linkedin.com/jobs/view/1404185031/")

}
