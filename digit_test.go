package bench_test

import (
	"testing"
	"unicode"

	"bench"
)

// goos: darwin
// goarch: arm64
// BenchmarkIsDigitUnicode-8       74610045                90.72 ns/op           96 B/op          0 allocs/op
// BenchmarkIsDigitMap-8           23898244               214.9 ns/op            98 B/op          0 allocs/op

func BenchmarkIsDigitUnicode(b *testing.B) {
	var z []rune
	s := "%^80      hhhhh20&&&&nd"

	for i := 0; i < b.N; i++ {
		for _, v := range s {
			if unicode.IsDigit(v) {
				z = append(z, v)
			}
		}
	}
}

func BenchmarkIsDigitMap(b *testing.B) {
	var z []rune
	s := "%^80      hhhhh20&&&&nd"

	for i := 0; i < b.N; i++ {
		for _, v := range s {
			if bench.IsDigit(v) {
				z = append(z, v)
			}
		}
	}
}
