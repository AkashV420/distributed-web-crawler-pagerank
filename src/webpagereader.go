package main

import (
	"fmt"

	"golang.org/x/net/html"
)

// Extract retrieves the information from the webpage body
func extractContent(webpageURL Urlnode, crawedLinksChannel chan Urlnode) {

	fmt.Println("Trying to extract ", webpageURL.Url, "with depth: ", webpageURL.Depth)
	response, success := ConnectToWebsite(webpageURL.Url)

	if !success {
		fmt.Println("Received error while connecting to website: ", webpageURL.Url, "and depth: ", webpageURL.Depth)
		return
	}

	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()

		if isAnchorTag(tokenType, token) {
			cl, ok := extractLinksFromToken(token, webpageURL.Url)

			if ok {
				go func() {
					crawedLinksChannel <- Urlnode{cl, webpageURL.Depth - 1, webpageURL.Url}
				}()
			}
		}
	}
}
