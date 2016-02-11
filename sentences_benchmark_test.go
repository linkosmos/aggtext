package aggtext

import "testing"

func BenchmarkSentences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sentencesTests {
			Sentences(test.input)
		}
	}
}
