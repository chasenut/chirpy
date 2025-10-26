package main

import "strings"

func replaceAllWordsCaseInsensitive(in string, filters []string, replacement string) string {
	if len(filters) == 0 {
		return in
	}

	loweredFiltersSet := make(map[string]struct{})
	for _, filter := range filters {
		loweredFiltersSet[filter] = struct{}{}
	}

	words := strings.Split(in, " ")
	cleaned := []string{}
	for _, word := range words {
		lowered := strings.ToLower(word)
		if _, ok := loweredFiltersSet[lowered]; ok {
			cleaned = append(cleaned, replacement)
			continue
		}
		cleaned = append(cleaned, word)
	}
	return strings.Join(cleaned, " ")
}
