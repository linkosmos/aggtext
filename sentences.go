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

func isSentenceEnd(input ...rune) bool {
	return len(input) == 2 && (input[0] == ' ') && unicode.IsUpper(input[1])
}

func fields(text string) (output []string) {
	// Reduce multi whitespace to one
	text = MultiWhitespace.ReplaceAllString(text, SPACECHAR)

	var buffer []rune

	runes := []rune(text)

	foundN := 0
	fieldStart := -1 // Set to -1 when looking for start of field.
	for index, rune := range runes {

		if rune == '.' || rune == '!' || rune == '?' {

			// Ignoring trailing sentence ending e.g.: !!!, ...., ??
			if foundN == 1 {
				continue
			}

			foundN = 1
			continue
		}

		if foundN == 1 || foundN == 2 {
			foundN++
			buffer = append(buffer, rune)
		} else if foundN == 3 {
			if isSentenceEnd(buffer...) {
				// -2 due to that we are at 2nd position and
				// split should happen on space
				// e.g. ". A"
				// fmt.Println(
				// 	string(text[index-3]),
				// 	fmt.Sprintf("[%s]", string(text[index-2])),
				// 	string(text[index-1]),
				// )
				// -2 ignore whitespace
				// -3 ignore separation symbol: ., !, ?

				if fieldStart == -1 {
					output = append(output, string(runes[:index-2]))
				} else {
					output = append(output, string(runes[fieldStart:index-2]))
				}
				// +1 ignoring whitespace
				fieldStart = index - 1
			}

			// Reset
			foundN = 0
			buffer = nil
		}
	}
	if fieldStart >= 0 { // Last field might end at EOF.
		output = append(output, string(runes[fieldStart:]))
	} else {
		output = append(output, text)
	}
	return output
}

// Sentences - splits unicode text into Sentences
func Sentences(text ...string) (output []string) {
	// var buffer []string
	for _, txt := range text {

		output = append(output, fields(txt)...)

		// // Remove surrounding whitespace
		// txt = SurroundingWhitespapce.ReplaceAllString(txt, EMPTYCHAR)
		//
		// // Reduce multi whitespace to one
		// txt = MultiWhitespace.ReplaceAllString(txt, SPACECHAR)
		//
		// for _, word := range Sentence.Split(txt, -1) {
		// 	// Skip empty
		// 	if len(word) == 0 {
		// 		continue
		// 	}
		//
		// 	if size := len(buffer); size > 0 {
		//
		// 		// If word is lower it shouldn't be matched
		// 		if isLower(word) && size > 0 {
		// 			// In order to add original splitting symbol ./!/?
		// 			// need to track history or index
		// 			buffer[size-1] += ". " + word
		// 			continue
		// 		}
		// 	}
		//
		// 	buffer = append(buffer, word)
		// }
		//
		// output = append(output, buffer...)
	}
	return output
}
