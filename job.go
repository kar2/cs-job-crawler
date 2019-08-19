package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Job ...
type Job struct {
	company     string
	description string
	link        string
}

func getJobFromLink(link string) Job {
	response, err := http.Get(link)
	if err != nil {
		log.Fatal("Can't get response from link: " + link)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP content.")
	}
	var job Job
	job.link = link
	job.description = doc.Find("h1").Text()
	job.company = doc.Find("a").First().Text()
	fmt.Println(job.company)
	return job
}
func process(linkMap map[int]string) map[int]Job {
	jobs := make(map[int]Job)
	for id, link := range linkMap {
		jobs[id] = getJobFromLink(link)
	}
	return jobs
}
