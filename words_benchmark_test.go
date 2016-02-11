package aggtext

import "testing"

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range wordsTests {
			Words(test.input)
		}
	}
}
