package main

import (
	"fmt"
	"sync"
)

type safeCache struct {
	v   map[string]string
	mux sync.Mutex
}

func (c *safeCache) isCached(key string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.v[key]
	return ok
}

func (c *safeCache) setValue(key, value string) {
	c.mux.Lock()
	c.v[key] = value
	c.mux.Unlock()
}

var cache = safeCache{
	v: make(map[string]string),
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	cache.mux.Lock()
	if _, ok := cache.v[url]; ok {
		cache.mux.Unlock()
		return
	}
	cache.v[url] = "fetching" // TODO: error handling
	cache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)

	cache.setValue(url, body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}

	for i, u := range urls {
		_ = i
		_ = u
		<-done
	}

	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
