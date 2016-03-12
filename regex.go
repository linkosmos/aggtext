package aggtext

import "regexp"

// -
const (
	EMPTYCHAR = ""
	SPACECHAR = " "
)

// -
var (
	// Matches unicode punctuation
	// punctuation (== [!-/:-@[-`{-~])
	// https://github.com/google/re2/blob/master/doc/syntax.txt#L214
	Punctuation = regexp.MustCompile(`\p{P}|\|`)

	// Matches all multiple whitespace
	// whitespace (== [\t\n\v\f\r ])
	// blank (== [\t ])
	MultiWhitespace = regexp.MustCompile(`\s{2,}`)
)
