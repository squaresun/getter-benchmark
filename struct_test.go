package test

import (
	"testing"
)

func BenchmarkServeStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serveStruct()
	}
}
