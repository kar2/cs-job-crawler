package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Job ...
type Job struct {
	company string
	role    string
	link    string
}

func constructJobFromLink(link string) Job {

	// Output Job struct
	var job Job
	job.company = ""
	job.role = ""
	job.link = link

	// Get http from link
	response, err := http.Get(link)
	if err != nil {
		log.Fatal("Can't get response from link: " + link)
	}
	defer response.Body.Close()

	// Construct document from http
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP content.")
	}

	// Parse company and job info
	sep := " "
	metaInfo := strings.Trim(strings.Split(doc.Find("title").Text(), "|")[0], sep)
	tokenizedInfo := strings.Fields(metaInfo)

	// Get company name
	company := ""
	for _, token := range tokenizedInfo {
		if token == "hiring" {
			break
		} else {
			company = company + token + " "
		}
	}
	job.company = company[0 : len(company)-1]

	// Get role
	foundRole := false
	role := ""
	for _, token := range tokenizedInfo {
		if foundRole {
			if token == "in" {
				break
			} else {
				role = role + token + " "
			}
		} else if token == "hiring" {
			foundRole = true
		}
	}
	job.role = role[0 : len(role)-1]

	return job
}

func process(linkMap map[int]string) map[int]Job {

	// Map from Linkedin page ID to Job struct
	jobs := make(map[int]Job)
	for id, link := range linkMap {
		jobs[id] = constructJobFromLink(link)
	}
	fmt.Println("Successfully processed " + strconv.Itoa(maxPages) + " jobs.")
	return jobs
}
