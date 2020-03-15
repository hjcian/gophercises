package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/hjcian/gophercises/link"
)

const xmlNameSpace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	toXML := urlset{
		Xmlns: xmlNameSpace,
	}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	xmlEnc := xml.NewEncoder(os.Stdout)
	xmlEnc.Indent("", "  ")
	if err := xmlEnc.Encode(toXML); err != nil {
		panic(err)
	}
	fmt.Println()
}

// Optimization trick: https://dave.cheney.net/2014/03/25/the-empty-struct
type emptyStruct struct{}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]emptyStruct)
	var q map[string]emptyStruct
	nq := map[string]emptyStruct{
		urlStr: emptyStruct{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]emptyStruct)
		if len(q) == 0 {
			// break the loop before the i doesn't meet the maxDepth
			break
		}
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = emptyStruct{}
			for _, link := range get(url) {
				nq[link] = emptyStruct{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		// panic(err)
		// just return empty slice of string is OK in this scenario
		return []string{}
	}
	defer resp.Body.Close()
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(body io.Reader, base string) []string {
	links, _ := link.Parse(body)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(link string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(base string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, base)
	}
}
