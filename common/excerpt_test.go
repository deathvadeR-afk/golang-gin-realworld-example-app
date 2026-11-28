package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// excerptLongBody is a 400-character article body made of repeated 9-letter words.
func excerptLongBody() string {
	return strings.Repeat("abcdefghi ", 40)
}

func TestExcerpt(t *testing.T) {
	t.Run("[QP-SMGIN-7-1] collapses whitespace and newlines and trims the ends", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Equal("Hello world", Excerpt("  Hello\n\nworld  "))
	})

	t.Run("[QP-SMGIN-7-2] returns a short body unchanged with no ellipsis", func(t *testing.T) {
		asserts := assert.New(t)
		excerpt := Excerpt("Short article body")
		asserts.Equal("Short article body", excerpt)
		asserts.NotContains(excerpt, "...")
	})

	t.Run("[QP-SMGIN-7-3] caps a long repeated-word body at 160 characters ending in an ellipsis", func(t *testing.T) {
		asserts := assert.New(t)
		excerpt := Excerpt(excerptLongBody())
		asserts.LessOrEqual(len(excerpt), 160)
		asserts.True(strings.HasSuffix(excerpt, "..."))
	})

	t.Run("[QP-SMGIN-7-4] ends the excerpt at the last whole word that fits", func(t *testing.T) {
		asserts := assert.New(t)
		excerpt := Excerpt(excerptLongBody())
		kept := strings.TrimSuffix(excerpt, "...")
		for _, word := range strings.Fields(kept) {
			asserts.Equal("abcdefghi", word)
		}
		asserts.False(strings.HasSuffix(kept, " "))
		// One more 9-letter word and its space would add 10 characters to the excerpt.
		asserts.Greater(len(excerpt)+10, 160)
	})

	t.Run("[QP-SMGIN-7-5] runs with no database and no running server", func(t *testing.T) {
		asserts := assert.New(t)
		asserts.Equal("Hello world", Excerpt("  Hello\n\nworld  "))
	})
}
