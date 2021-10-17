package main

import (
	"fmt"
	"github.com/shashankdev81/web-crawler/crawler"
)

func main() {
	var crawledUrls = crawler.Crawl("https://yahoo.com")
	for url := range crawledUrls {
		fmt.Println(url)
	}

}
