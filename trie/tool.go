package trie

import "strings"

// Part split the URL
func splitUrl(path string) []string {
	var parts = make([]string, 0, 2)
	for _, val := range strings.Split(path, "/") {
		if val != "/" && val != "" {
			parts = append(parts, val)
		}
	}
	return parts
}

func join(parts []string) string {
	if len(parts) == 0 {
		return "/"
	}
	return strings.Join(append([]string{""}, parts...), "/")
}

func format(path string) string {
	return join(splitUrl(path))
}
