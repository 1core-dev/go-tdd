package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
