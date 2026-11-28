package articles

import "strings"

// excerptMaxLength is the longest excerpt, in characters, returned unchanged.
const excerptMaxLength = 160

// excerptEllipsis is appended to an excerpt that has been shortened.
const excerptEllipsis = "..."

// excerpt builds a short preview of an article body. Whitespace is collapsed
// into single spaces and trimmed. A collapsed body longer than excerptMaxLength
// is cut to the longest run of whole words that fits once excerptEllipsis is
// appended, and the ellipsis is added. The body itself is not changed.
func excerpt(body string) string {
	collapsed := strings.Join(strings.Fields(body), " ")
	if len(collapsed) <= excerptMaxLength {
		return collapsed
	}

	limit := excerptMaxLength - len(excerptEllipsis)
	prefix := ""
	for _, word := range strings.Fields(collapsed) {
		candidate := word
		if prefix != "" {
			candidate = prefix + " " + word
		}
		if len(candidate) > limit {
			break
		}
		prefix = candidate
	}
	return prefix + excerptEllipsis
}
