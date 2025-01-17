package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int){
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil{
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return 
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil{
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return 
	}
	
	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname(){
		return 
	}
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}
	
	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return
		
	}

	pages[normalizedCurrentURL] = 1
	
	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil{
		fmt.Printf("Error - getHTML: %v", err)
		return
	}
	
	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}
	for _, url := range urls{
		crawlPage(rawBaseURL, url, pages)
	}
}