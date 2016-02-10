package aggtext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var punctuationTests = []struct {
	input    string
	expected string
}{
	{
		"-_-",
		"",
	},
	{
		"(1)«",
		"1",
	},
	{
		"¶ Para",
		" Para",
	},
}

func TestPunctuation(t *testing.T) {
	for _, test := range punctuationTests {
		got := Punctuation.ReplaceAllString(test.input, "")

		assert.Equal(t, test.expected, got)
	}
}

var surroundingWhitespapceTests = []struct {
	input    string
	expected string
}{
	{
		" Su ",
		"Su",
	},
	{
		" S u ",
		"S u",
	},
	{
		"		Tab		",
		"Tab",
	},
	{
		`
	
		New Line
		
		`,
		"New Line",
	},
}

var multiWhitespaceTests = []struct {
	input    string
	expected string
}{
	{
		"Multi  whitespace to one space",
		"Multi whitespace to one space",
	},
	{
		"Multi		tab to one space",
		"Multi tab to one space",
	},
	{
		"Multi		tab to one space",
		"Multi tab to one space",
	},
	{
		"    whitespace     ",
		" whitespace ",
	},
}

func TestMultiWhitespace(t *testing.T) {
	for _, test := range multiWhitespaceTests {
		got := MultiWhitespace.ReplaceAllString(test.input, " ")

		assert.Equal(t, test.expected, got)
	}
}

func TestSurroundingWhitespapce(t *testing.T) {
	for _, test := range surroundingWhitespapceTests {
		got := SurroundingWhitespapce.ReplaceAllString(test.input, "")

		assert.Equal(t, test.expected, got)
	}
}
