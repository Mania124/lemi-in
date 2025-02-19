package api

import "strings"

// Split splits a string by the '-' delimiter and returns a slice of substrings.
func Split(s string) []string {
	var result []string
	var builder strings.Builder

	for _, c := range s {
		if c == '-' {
			result = append(result, builder.String())
			builder.Reset()
		} else {
			builder.WriteRune(c)
		}
	}

	// Append the last segment
	result = append(result, builder.String())
	return result
}
