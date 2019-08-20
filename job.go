package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	fmt.Println("Processing jobs...")
	// Map from Linkedin page ID to Job struct
	jobs := make(map[int]Job)
	for id, link := range linkMap {
		jobs[id] = constructJobFromLink(link)
	}
	fmt.Println("Successfully processed " + strconv.Itoa(len(linkMap)) + " jobs.")
	return jobs
}

func exportAsTSV(jobMap map[int]Job) int {
	fmt.Println()
	fmt.Println("Exporting jobs as TSV...")

	sep := "\t"

	// Create output file
	tsv, err := os.Create("./jobs.tsv")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer tsv.Close()

	// Write header
	tsv.WriteString("Company" + sep + "Role" + sep + "Link")
	tsv.WriteString("\n")

	// Write to tsv
	for _, job := range jobMap {
		tsv.WriteString(job.company + sep + job.role + sep + job.link)
		tsv.WriteString("\n")
	}
	tsv.Sync()

	fmt.Println("Finished exporting.")
	return 0
}
