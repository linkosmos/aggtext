package aggtext

import "strings"

// Words - removes punctuation and splits text into words
func Words(text ...string) (output []string) {
	return WordsFunc(func(input string) string {
		return Punctuation.ReplaceAllString(input, EMPTYCHAR)
	}, text...)
}

// WordsFunc - takes function to aplly before splitting given text into words
func WordsFunc(funcToApply func(string) string, text ...string) (output []string) {
	for _, txt := range text {
		output = append(output, strings.Fields(funcToApply(txt))...)
	}
	return output
}
