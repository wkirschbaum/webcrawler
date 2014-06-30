package crawler

import (
	"fmt"
	"sync"

	"github.com/unboxed/webcrawler/fetcher"
)

type Crawler struct {
	Fetching    chan fetcher.FetchedResult
	Done        chan bool
	Depth       int
	Concurrency int
}

func (crawler *Crawler) Crawl(url string, f fetcher.Fetcher) {
	m := map[string]bool{url: true}

	limitChan := make(chan bool, crawler.Concurrency)

	var mx sync.Mutex
	var wg sync.WaitGroup
	var c2 func(string, int)

	c2 = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}

		limitChan <- true
		fetchedResult, err := f.Fetch(url)
		<-limitChan

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fetchedResult.Level = (crawler.Depth - depth)
		crawler.Fetching <- *fetchedResult

		mx.Lock()
		for _, u := range fetchedResult.Urls {
			if !m[u] {
				m[u] = true
				wg.Add(1)
				go c2(u, depth-1)
			}
		}
		mx.Unlock()
	}

	wg.Add(1)
	c2(url, crawler.Depth)
	wg.Wait()
	crawler.Done <- true
}
