package aggtext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var alphaTests = []struct {
	input    string
	expected string
}{
	{
		"awkdjk 2  10. . .2$--- !!",
		"awkdjk 2 10 2$",
	},
	{
		"awkd            jk 2  10. . .2--- !!",
		"awkd jk 2 10 2",
	},
	{
		"awkd            jk 2  10. . .2––--- !!",
		"awkd jk 2 10 2",
	},
}

func TestAlpha(t *testing.T) {
	for _, test := range alphaTests {
		got := Alpha(test.input)

		assert.Equal(t, test.expected, got)
	}
}
