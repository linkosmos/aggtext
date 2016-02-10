package aggtext

// Alpha - remove punctuation, reduce multiple whitespace into 1 and
// remove surrounding whitespace
func Alpha(s string) string {
	s = Punctuation.ReplaceAllString(s, EMPTYCHAR)
	s = MultiWhitespace.ReplaceAllString(s, SPACECHAR)
	s = SurroundingWhitespapce.ReplaceAllString(s, EMPTYCHAR)
	return s
}
