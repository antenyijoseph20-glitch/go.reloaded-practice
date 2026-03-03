package main

import (
	"strings"
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Hex conversion",
			input:    "1E (hex) files",
			expected: "30 files",
		},
		{
			name:     "Binary conversion",
			input:    "10 (bin) years",
			expected: "2 years",
		},
		{
			name:     "Uppercase single word",
			input:    "ready set go (up)",
			expected: "ready set GO",
		},
		{
			name:     "Capitalize multiple words",
			input:    "it was the age of foolishness (cap, 6)",
			expected: "It Was The Age Of Foolishness",
		},
		{
			name:     "Article a to an",
			input:    "There was a amazing apple and a hour",
			expected: "There was an amazing apple and an hour",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words := strings.Fields(tt.input)
			result := Transform(words)
			got := strings.Join(result, " ")
			if got != tt.expected {
				t.Errorf("Transform() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic punctuation spacing",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "Ellipsis handling",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Single quotes wrapping",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPunctuation(tt.input)
			if got != tt.expected {
				t.Errorf("FormatPunctuation() = %v, want %v", got, tt.expected)
			}
		})
	}
}
