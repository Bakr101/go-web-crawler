package main

import (
	"fmt"
	"os"
)

func main(){
	fullCLIArgs := os.Args
	commands := fullCLIArgs[1:]
	if len(commands) < 1{
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(commands) > 1{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(commands) == 1{
		fmt.Printf("starting crawl of: %s\n", commands[0])
		html, err := getHTML(commands[0])
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(html)
	}
}