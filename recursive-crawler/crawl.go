package crawler

import (
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}
var crawledSet = make(map[string]bool)
var lock = sync.RWMutex{}

func Crawl(url string) []string {
	var results = []string{}
	wg.Add(1)
	traverse(url)
	wg.Wait()
	for k, _ := range crawledSet {
		results = append(results, k)
	}
	return results
}

func traverse(url string) {
	lock.RLock()
	_, ok := crawledSet[url]
	lock.RUnlock()
	if ok {
		wg.Done()
		return
	}
	lock.Lock()
	crawledSet[url] = true
	lock.Unlock()
	var fetched = fetch(url)
	for _, crawled := range fetched {
		wg.Add(1)
		go traverse(crawled)
	}
	wg.Done()
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

	var ind = rand.Intn(10)
	return urls[ind : ind+2]
}
