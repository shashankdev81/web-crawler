package main

import (
	"fmt"
	"github.com/shashankdev81/web-crawler/recursive_crawler"
)

func main() {
	var crawledUrls = recursive_crawler.Crawl("https://yahoo.com")
	for _, url := range crawledUrls {
		fmt.Println(url)
	}

}
