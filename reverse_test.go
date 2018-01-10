package strtools

import (
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reverse("AaaaaaaaaaBBBBBBBBBccccccccccDDDDDDDD")
	}
}
