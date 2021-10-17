package main

import (
	"fmt"
	"github.com/shashankdev81/web-crawler/CrawlWithChannels"
	"github.com/shashankdev81/web-crawler/CrawlWithRecursion"
)

func main() {
	var crawledUrls = recursive_crawler.Crawl("https://yahoo.com")
	for _, url := range crawledUrls {
		fmt.Println(url)
	}

	var crawledUrls2 = crawler_with_channel.Crawl("https://yahoo.com")
	for _, url := range crawledUrls2 {
		fmt.Println(url)
	}

}
