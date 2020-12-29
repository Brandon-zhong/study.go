package benchmark_demo

import (
	"strconv"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

func BenchmarkDemo(b *testing.B) {
	str := ""
	for n := 0; n < b.N; n++ {
		str += strconv.Itoa(n)
	}
}
