package aggtext

import "strings"

// Words - removes punctuation and splits text into words
func Words(text ...string) []string {
	return WordsFunc(Alpha, text...)
}

// WordsFunc - takes function to aplly before splitting given text into words
func WordsFunc(funcToApply func(string) string, text ...string) (output []string) {
	for _, txt := range text {
		for _, word := range strings.SplitN(funcToApply(txt), SPACECHAR, -1) {
			if len(word) > 0 {
				output = append(output, word)
			}
		}
	}
	return output
}
