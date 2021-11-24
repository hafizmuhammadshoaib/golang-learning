package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlingImpl struct {
	mu       sync.Mutex
	wg       sync.WaitGroup
	readUrls map[string]bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *CrawlingImpl) Crawl(url string, depth int, fetcher Fetcher) {
	defer c.mu.Unlock()
	defer c.wg.Done()
	if depth <= 0 {
		return
	}
	c.mu.Lock()
	if !c.readUrls[url] {
		body, urls, err := fetcher.Fetch(url)
		c.readUrls[url] = true

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			{
				c.wg.Add(1)
				go c.Crawl(u, depth-1, fetcher)
			}
		}
	}

	return
}

func main() {
	crawling := CrawlingImpl{readUrls: make(map[string]bool)}
	crawling.wg.Add(1)
	go crawling.Crawl("https://golang.org/", 4, fetcher)
	crawling.wg.Wait()
	fmt.Print(crawling.readUrls)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
