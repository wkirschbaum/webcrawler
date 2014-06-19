package crawler

import (
	"sync"

	"github.com/wkirschbaum/webcrawler/fetcher"
)

func Crawl(url string, depth int, f fetcher.Fetcher) []fetcher.FetchedResult {
	m := map[string]fetcher.FetchedResult{}
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
		mx.Lock()
		for _, u := range fetchedResult.Urls {
			if _, ok := m[u]; !ok {
				m[u] = *fetchedResult
				wg.Add(1)
				go c2(u, depth-1)
			}
		}
		mx.Unlock()
	}
	wg.Add(1)
	c2(url, depth)
	wg.Wait()

	v := make([]fetcher.FetchedResult, len(m), len(m))
	idx := 0
	for _, value := range m {
		v[idx] = value
		idx++
	}
	return v
}
