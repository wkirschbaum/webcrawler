package fetcher

import (
	"io/ioutil"
	"net/http"

	"github.com/unboxed/webcrawler/parser"
)

type Fetcher interface {
	Fetch(url string) (result *FetchedResult, err error)
}

type liveFetcher struct {
	BaseUrl string
}

type FetchedResult struct {
	Url    string
	Body   string
	Urls   []string
	Status int
}

func (f liveFetcher) Fetch(url string) (*FetchedResult, error) {
	body := ""
	statusCode := 0

	resp, err := http.Get(url)

	if err == nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		body = string(bodyBytes)
		statusCode = resp.StatusCode
	}

	p := parser.Parser{BaseUrl: f.BaseUrl}
	links := p.ParseLinks(body)

	return &FetchedResult{Url: url, Body: body, Urls: links, Status: statusCode}, err
}

func MakeFetcher(baseUrl string) Fetcher {
	return &liveFetcher{BaseUrl: baseUrl}
}
