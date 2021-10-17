package main

import (
	"fmt"
	"github.com/shashankdev81/web-crawler/CrawlWithRecursion"
)

func main() {
	var crawledUrls = crawler.Crawl("https://yahoo.com")
	for _, url := range crawledUrls {
		fmt.Println(url)
	}

}
