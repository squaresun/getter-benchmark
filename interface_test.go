package benchmark

import "testing"

func BenchmarkServeInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serveInterface()
	}
}
