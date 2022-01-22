package goruntine

import "testing"

func BenchmarkCreateSConfigUseRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSConfigUseRWMutex()
	}
}

func BenchmarkCreateSConfigUseAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSConfigUseAtomic()
	}
}
