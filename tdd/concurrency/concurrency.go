package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	url    string
	result bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan Result)

	for _, url := range urls {
		// results[url] = wc(url)
		go func(url string) {
			resultChannel <- Result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.result
	}
	return results
}
