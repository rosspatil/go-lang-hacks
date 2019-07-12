package main

import (
	"testing"
)

func Benchmark_simpleBufferOperation(b *testing.B) {
	for index := 0; index < b.N; index++ {
		hitmesimple()
	}
}

func Benchmark_pooledBufferOperation(b *testing.B) {
	for index := 0; index < b.N; index++ {
		hitmepooled()
	}
}
