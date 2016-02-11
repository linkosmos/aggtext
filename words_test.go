package aggtext

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var wordsTests = []struct {
	input    string
	expected []string
}{
	{
		"It be distracted by the    readable content of a page when looking at its layout.",
		[]string{
			"It", "be", "distracted", "by", "the",
			"readable", "content", "of", "a", "page", "when", "looking", "at", "its", "layout",
		},
	},
	{
		"This isin't simple",
		[]string{
			"This", "isint", "simple",
		},
	},
	{
		`It be   distracted !by! 
		the readable content ,of, a 
				page 2292 when		looking at its layout.`,
		[]string{
			"It", "be", "distracted", "by", "the",
			"readable", "content", "of", "a", "page", "2292",
			"when", "looking", "at", "its", "layout",
		},
	},
	{
		"aw(@!())",
		[]string{
			"aw",
		},
	},
}

func TestWords(t *testing.T) {
	for _, test := range wordsTests {
		got := Words(test.input)

		assert.Equal(t, test.expected, got)
	}
}

var wordsFuncTests = []struct {
	input    string
	f        func(string) string
	expected []string
}{
	{
		"It be distracted by",
		func(s string) string {
			return s
		},
		[]string{
			"It", "be", "distracted", "by",
		},
	},
	{
		"This isin't simple password",
		func(s string) string {
			return strings.Replace(s, "password", "******", -1)
		},
		[]string{
			"This", "isin't", "simple", "******",
		},
	},
}

func TestWordsFunc(t *testing.T) {
	for _, test := range wordsFuncTests {
		got := WordsFunc(test.f, test.input)

		assert.Equal(t, test.expected, got)
	}
}
