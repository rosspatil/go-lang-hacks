package main

import "testing"

func BenchmarkSetInGabs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetInGabs()
	}
}

func BenchmarkGjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetInGjson()
	}
}
