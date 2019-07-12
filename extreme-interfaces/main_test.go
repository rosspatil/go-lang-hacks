package main

import (
	"fmt"
	"testing"
)

func Benchmark_useInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(useInterface())
	}
}

func Benchmark_useStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(useStruct())
	}
}
