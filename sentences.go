package aggtext

import "unicode"

// Faster version than to compare whole sentence to strings.IsLower
// OR strings.ToUpper OR cast into []rune slice
func isLower(sentence string) bool {
	for _, letters := range sentence {
		return unicode.IsLower(letters)
	}
	return false
}

// Sentences - splits unicode text into Sentences
func Sentences(text ...string) (output []string) {
	var buffer []string
	for _, txt := range text {

		// Remove surrounding whitespace
		txt = SurroundingWhitespapce.ReplaceAllString(txt, EMPTYCHAR)

		// Reduce multi whitespace to one
		txt = MultiWhitespace.ReplaceAllString(txt, SPACECHAR)

		for _, word := range Sentence.Split(txt, -1) {
			// Skip empty
			if len(word) == 0 {
				continue
			}

			if size := len(buffer); size > 0 {

				// If word is lower it shouldn't be matched
				if isLower(word) && size > 0 {
					// In order to add original splitting symbol ./!/?
					// need to track history or index
					buffer[size-1] += ". " + word
					continue
				}
			}

			buffer = append(buffer, word)
		}

		output = append(output, buffer...)
	}
	return output
}
