package main

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single word",
			input:    "Hello",
			expected: 1,
		},
		{
			name:     "Multiple words",
			input:    "Это тестовая строка",
			expected: 3,
		},
		{
			name:     "Extra spaces",
			input:    "  Много   пробелов  между  словами  ",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wordCount := countWords(tt.input)
			if wordCount != tt.expected {
				t.Errorf("countWords() = %v, expected %v", wordCount, tt.expected)
			}
		})
	}
}
