package articles

import (
	"strings"
	"unicode/utf8"
)

// excerptMaxLength is the longest excerpt in characters, ellipsis included.
const excerptMaxLength = 160

// excerptEllipsis is appended to an excerpt that cuts the body short.
const excerptEllipsis = "..."

// excerpt returns a plain-text summary of an article body. Every run of
// whitespace becomes one space and both ends are trimmed. A text of at most
// excerptMaxLength characters is returned as it is. A longer text is cut after
// its last whole word that still leaves room for the ellipsis. The input string
// is not changed.
func excerpt(body string) string {
	text := strings.Join(strings.Fields(body), " ")
	if utf8.RuneCountInString(text) <= excerptMaxLength {
		return text
	}
	runes := []rune(text)
	limit := excerptMaxLength - utf8.RuneCountInString(excerptEllipsis)
	// cut is the position of the last space at or before limit: the text before
	// it is a run of whole words.
	cut := 0
	for i := 0; i <= limit && i < len(runes); i++ {
		if runes[i] == ' ' {
			cut = i
		}
	}
	return string(runes[:cut]) + excerptEllipsis
}
