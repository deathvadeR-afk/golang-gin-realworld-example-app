package common

import (
	"strings"
	"unicode/utf8"
)

// excerptLimit is the longest excerpt returned, counting the trailing ellipsis.
const excerptLimit = 160

// excerptEllipsis is appended to an excerpt that was cut short.
const excerptEllipsis = "..."

// Excerpt builds a plain-text excerpt of an article body. It collapses every
// run of whitespace and newlines to one space and trims both ends. A body of
// 160 characters or fewer is returned unchanged; a longer one is cut at the
// last whole word that, with "..." appended, stays within 160 characters.
// The input string is not modified.
func Excerpt(body string) string {
	words := strings.Fields(body)
	text := strings.Join(words, " ")
	if utf8.RuneCountInString(text) <= excerptLimit {
		return text
	}

	kept := ""
	for _, word := range words {
		candidate := word
		if kept != "" {
			candidate = kept + " " + word
		}
		if utf8.RuneCountInString(candidate)+utf8.RuneCountInString(excerptEllipsis) > excerptLimit {
			break
		}
		kept = candidate
	}
	return kept + excerptEllipsis
}
