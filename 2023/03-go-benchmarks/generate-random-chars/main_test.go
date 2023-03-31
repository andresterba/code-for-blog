package main

import (
	"testing"
)

func Benchmark_algStringCasting(b *testing.B) {
	for n := 0; n < b.N; n++ {
		algStringCasting()
	}
}

func Benchmark_algArrayLookup(b *testing.B) {
	for n := 0; n < b.N; n++ {
		algArrayLookup()
	}
}
