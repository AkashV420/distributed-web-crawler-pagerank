package main

import "fmt"
import "sync"
// ProcessCrawledLinks reads the crawled links and adds unique links to sitesChannel
func ProcessCrawledLinks(wg *sync.WaitGroup, sitesChannel chan Urlnode, crawedLinksChannel chan Urlnode, pendingCountChannel chan int) {
	foundUrls := make(map[string]bool)
	g := NewGraph()
	
	for cl := range crawedLinksChannel {
		g.AddEdge(NodeID(cl.Parenturl), NodeID(cl.Url))

		if cl.Depth >= 1 && !foundUrls[cl.Url] {
			foundUrls[cl.Url] = true
			pendingCountChannel <- 1
			sitesChannel <- cl
		}
	}
	fmt.Println("edges", g.GetAmountOfEdges())
	fmt.Println("nodes", g.GetAmountOfNodes())
	pr := NewPageRank(g)
	fmt.Println("here 1")
	pr.CalcPageRank()
	fmt.Println("here 2")
	pr.OrderResults()
	fmt.Println("Max to Min")
	for i, k := range pr.GetMaxToMinOrder() {
		fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
		if i > 1000 {
			break
		}
	}
	pr.ExportToCSV("final-rank.csv")
	fmt.Println("Done processing links")
	wg.Done()
}
