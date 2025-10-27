package main

import "strings"

func replaceAllWordsCaseInsensitive(in string, filters []string, replacement string) string {
	if len(filters) == 0 {
		return in
	}

	filtersSet := make(map[string]struct{})
	for _, filter := range filters {
		filtersSet[filter] = struct{}{}
	}

	words := strings.Split(in, " ")
	for i, word := range words {
		lowered := strings.ToLower(word)
		if _, ok := filtersSet[lowered]; ok {
			words[i] = replacement
		}
	}
	return strings.Join(words, " ")
}
