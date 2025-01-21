package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string){
	pagesSlice := sortPages(pages)
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	for _, page := range pagesSlice {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
	fmt.Println("=============================")
}
type page struct {
	url string
	count int
}
func sortPages(pages map[string]int) []page{
	pagesSlice := []page{}
	for k,v := range pages{
		pagesSlice = append(pagesSlice, page{url: k, count: v})
	}
	sort.Slice(pagesSlice, func(i, j int) bool {
		return pagesSlice[i].count > pagesSlice[j].count
	})
	return pagesSlice
}