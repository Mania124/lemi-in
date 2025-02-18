package api

import (
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a-b-c", []string{"a", "b", "c"}},
		{"hello-world", []string{"hello", "world"}},
		{"a--b", []string{"a", "", "b"}}, // Handling consecutive delimiters
		{"singleword", []string{"singleword"}}, // No delimiter in the string
		{"", []string{""}}, // Empty string case
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := Split(test.input)
			if !equal(got, test.expected) {
				t.Errorf("Split(%q) = %v; want %v", test.input, got, test.expected)
			}
		})
	}
}

// Helper function to compare slices
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
