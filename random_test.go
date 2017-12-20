package strtools

import (
	"testing"
)

func benchmarkRandom(l int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Random(l)
	}
}

func BenchmarkRandom10(b *testing.B) { benchmarkRandom(10, b) }
func BenchmarkRandom15(b *testing.B) { benchmarkRandom(15, b) }
func BenchmarkRandom25(b *testing.B) { benchmarkRandom(25, b) }
