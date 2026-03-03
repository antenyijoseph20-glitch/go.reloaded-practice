package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Transform(words []string) []string {
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				val, _ := strconv.ParseInt(words[i-1], 16, 64)
				words[i-1] = fmt.Sprintf("%d", val)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(bin)":
			if i > 0 {
				val, _ := strconv.ParseInt(words[i-1], 2, 64)
				words[i-1] = fmt.Sprintf("%d", val)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up,", "(low,", "(cap,":
			tag := words[i]
			countStr := strings.Trim(words[i+1], "皮)")
			count, _ := strconv.Atoi(countStr)
			for j := 1; j <= count && (i-j) >= 0; j++ {
				if tag == "(up," {
					words[i-j] = strings.ToUpper(words[i-j])
				} else if tag == "(low," {
					words[i-j] = strings.ToLower(words[i-j])
				} else {
					words[i-j] = capitalize(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	// Handle 'a' vs 'an'
	for i := 0; i < len(words)-1; i++ {
		lowWord := strings.ToLower(words[i])
		if lowWord == "a" || lowWord == "an" {
			nextWord := strings.ToLower(words[i+1])
			if len(nextWord) > 0 && strings.ContainsAny(string(nextWord[0]), "aeiouh") {
				words[i] = matchCase(words[i], "an")
			} else {
				words[i] = matchCase(words[i], "a")
			}
		}
	}
	return words
}

func FormatPunctuation(text string) string {
	// Fix punctuation spacing: "word ," -> "word,"
	rePunc := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = rePunc.ReplaceAllString(text, "$1")

	// Ensure space after: "word,next" -> "word, next"
	reSpace := regexp.MustCompile(`([.,!?:;]+)([^s.,!?:;\s])`)
	text = reSpace.ReplaceAllString(text, "$1 $2")

	// Handle single quotes: ' word ' -> 'word'
	reQuote := regexp.MustCompile(`'\s+(.*?)\s+'`)
	text = reQuote.ReplaceAllString(text, "'$1'")

	return text
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func matchCase(original, target string) string {
	if len(original) > 0 && original[0] >= 'A' && original[0] <= 'Z' {
		return strings.ToUpper(string(target[0])) + target[1:]
	}
	return target
}
