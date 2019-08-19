package main

import (
	"github.com/gocolly/colly"
)

// Preference ...
type Preference struct {
	internship bool
	skills     []string
}

func getJobDesc(link string, c colly.Collector) {

}
func process(linkMap map[int]string, preferences Preference) {
	for id, link := range linkMap {

	}
}
