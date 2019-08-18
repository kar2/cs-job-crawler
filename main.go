package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Starting...")
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		if strings.Contains(link, "/jobs/view/") {
			fmt.Println("New job link: " + link)
		} else {
			return
		}

		c.Visit(e.Request.AbsoluteURL(link))

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
