package main

// MonitorCrawling looks for pending number of sites to crawl and closes the sitesQueue when there are no pending sites
func MonitorCrawling(sitesChannel chan Urlnode, crawedLinksChannel chan Urlnode, pendingCountChannel chan int) {

	count := 0

	for c := range pendingCountChannel {
		count += c
		if count == 0 {
			close(sitesChannel)
			close(crawedLinksChannel)
			close(pendingCountChannel)
		}
	}
}
