package main

import (
	"sync"
)

type Urlnode struct {
	Url string
	Depth int
	Parenturl string
}
// CrawlerMain is the main trigger for webcrawler
func main() {

	sitesChannel := make(chan Urlnode)
	crawedLinksChannel := make(chan Urlnode)
	pendingCountChannel := make(chan int)

	siteToCrawl := "https://www.linkedin.com"
	depth := 5

	go func() {
		crawedLinksChannel <- Urlnode{siteToCrawl, depth, ""}
	}()

	var wg sync.WaitGroup

	wg.Add(1)

	go ProcessCrawledLinks(&wg, sitesChannel, crawedLinksChannel, pendingCountChannel)
	go MonitorCrawling(sitesChannel, crawedLinksChannel, pendingCountChannel)

	var numCrawlerThreads = 50
	for i := 0; i < numCrawlerThreads; i++ {
		wg.Add(1)
		go CrawlWebpage(&wg, sitesChannel, crawedLinksChannel, pendingCountChannel)
	}

	wg.Wait()
}
