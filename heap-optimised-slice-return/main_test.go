package main

import "testing"

var data interface{}

func Benchmark_heapOptimised(b *testing.B) {
	for index := 0; index < b.N; index++ {
		// data = heapOptimised()
		data = normal()
	}
}
