package crawler

func crawl(url string) {

}

func fetch(url string) []string {
	var urls = []string{}

	for i := 0; i < 10; i++ {
		urls = append(urls, "https://finance.yahoo.com/")
		urls = append(urls, "https://news.yahoo.com/")
		urls = append(urls, "https://sports.yahoo.com/")
		urls = append(urls, "https://mobile.yahoo.com/")
		urls = append(urls, "https://shopping.yahoo.com/")
		urls = append(urls, "https://entertainment.yahoo.com/")

	}

	return urls

}
