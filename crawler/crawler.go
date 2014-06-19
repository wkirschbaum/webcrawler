package crawler

import (
	"sync"

	"github.com/wkirschbaum/webcrawler/fetcher"
)

func Crawl(url string, depth int, f fetcher.Fetcher, ch chan fetcher.FetchedResult, chDone chan bool) {
	m := map[string]bool{url: true}

	var mx sync.Mutex
	var wg sync.WaitGroup
	var c2 func(string, int)

	c2 = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}

		fetchedResult, err := f.Fetch(url)
		if err != nil {
			return
		}

		ch <- *fetchedResult

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
	c2(url, depth)
	wg.Wait()
	chDone <- true
}
