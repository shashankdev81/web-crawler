package main

import (
	"fmt"
	crawler_with_channel "github.com/shashankdev81/web-crawler/CrawlWithChannels"
)

func main() {
	//var crawledUrls = recursive_crawler.Crawl("https://yahoo.com")
	//for _, url := range crawledUrls {
	//	fmt.Println(url)
	//}
	//fmt.Println("Crawled urls=", len(crawledUrls))

	var crawledUrls2 = crawler_with_channel.Crawl("https://yahoo.com")
	for _, url := range crawledUrls2 {
		fmt.Println(url)
	}
	fmt.Println("Crawled urls=", len(crawledUrls2))

}
