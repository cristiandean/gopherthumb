package loaders

import (
	"strings"
)

// NormalizeURL return url normalized
func NormalizeURL(url string) string {
	if strings.HasPrefix(url, "/unsafe/") {
		return strings.Replace(url, "/unsafe/", "", 1)
	}
	return url
}
