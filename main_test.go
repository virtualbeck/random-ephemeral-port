package main

import (
	"testing"
)

func BenchmarkRandomPort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
