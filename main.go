package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	fullCLIArgs := os.Args
	commands := fullCLIArgs[1:]
	if len(commands) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(commands) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", commands[0])
	// html, err := getHTML(commands[0])
	// if err != nil{
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	
	if len(commands) == 3 {
		parsedBaseURL, err := url.Parse(commands[0])
		if err != nil {
			fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", commands[0], err)
			return
		}
		maxConcurrency, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Printf("Error: couldn't parse concurrency value: %v\n", err)
			return
		}
		maxPages, err := strconv.Atoi(commands[2])
		if err != nil {
			fmt.Printf("Error: couldn't parse max pages value: %v\n", err)
			return
		}
		
		cfg := Config{
			pages:              map[string]int{},
			baseURL:            parsedBaseURL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, maxConcurrency),
			wg:                 &sync.WaitGroup{},
			maxPages:           maxPages,
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(commands[0])
		cfg.wg.Wait()

		
		printReport(cfg.pages, commands[0])
	}
	if len(commands) == 1 {
		parsedBaseURL, err := url.Parse(commands[0])
		if err != nil {
			fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", commands[0], err)
			return
		}
		cfg := Config{
			pages:              map[string]int{},
			baseURL:            parsedBaseURL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, 5),
			wg:                 &sync.WaitGroup{},
			maxPages:           10,
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(commands[0])
		cfg.wg.Wait()

		
		printReport(cfg.pages, commands[0])
	}
	if len(commands) == 2 {
		parsedBaseURL, err := url.Parse(commands[0])
		if err != nil {
			fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", commands[0], err)
			return
		}
		maxConcurrency, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Printf("Error: couldn't parse concurrency value: %v\n", err)
			return
		}
		cfg := Config{
			pages:              map[string]int{},
			baseURL:            parsedBaseURL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, maxConcurrency),
			wg:                 &sync.WaitGroup{},
			maxPages:           10,
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(commands[0])
		cfg.wg.Wait()

		printReport(cfg.pages, commands[0])
	}

}
