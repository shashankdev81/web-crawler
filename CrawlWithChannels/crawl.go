package crawler_with_channel

import (
	"sync"
)

var crawledSet = make(map[string]bool)
var lock = sync.RWMutex{}

func Crawl(url string) []string {
	var results = []string{}
	var resultsChan = make(chan string, 10)
	go traverse(url, resultsChan)
	for crawled := range resultsChan {
		results = append(results, crawled)
	}
	return results

}

func traverse(url string, ret chan string) {
	defer channelClose(ret)
	lock.RLock()
	_, ok := crawledSet[url]
	lock.RUnlock()
	if ok {
		return
	}
	lock.Lock()
	crawledSet[url] = true
	lock.Unlock()
	ret <- url
	fetched := fetch(url)
	urlsChanArr := make([]chan string, len(fetched))
	for i, fetched := range fetched {
		urlsChanArr[i] = make(chan string, 10)
		go traverse(fetched, urlsChanArr[i])
	}

	for _, eachChannel := range urlsChanArr {
		for crawled := range eachChannel {
			ret <- crawled
		}
	}

	return

}

func channelClose(ret chan string) {
	close(ret)
}
func fetch(url string) []string {
	var urls = []string{}

	urls = append(urls, "https://finance.yahoo.com/")
	urls = append(urls, "https://news.yahoo.com/")
	urls = append(urls, "https://sports.yahoo.com/")
	urls = append(urls, "https://mobile.yahoo.com/")
	urls = append(urls, "https://shopping.yahoo.com/")
	urls = append(urls, "https://entertainment.yahoo.com/")
	urls = append(urls, "https://tech.yahoo.com/")
	urls = append(urls, "https://politics.yahoo.com/")
	urls = append(urls, "https://housing.yahoo.com/")
	urls = append(urls, "https://auto.yahoo.com/")
	urls = append(urls, "https://cosmo.yahoo.com/")

	//var ind = rand.Intn(10)
	//return urls[ind : ind+4]
	return urls
}
