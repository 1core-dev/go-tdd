package concurrency

import (
	"maps"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://unexisting.god"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://dave.cheney.net",
		"waat://unexisting.god",
	}

	want := map[string]bool{
		"http://google.com":      true,
		"http://dave.cheney.net": true,
		"waat://unexisting.god":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !maps.Equal(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}
