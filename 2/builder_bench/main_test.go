package main

import "testing"

const (
	str1 = "012345678"
	str2 = "012345678"
	str3 = "012345678"
)

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Concat(str1, str2, str3)
	}
}

func BenchmarkConcat10values(b *testing.B) {
	strs := []string{"012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678"}
	for i := 0; i < b.N; i++ {
		Concat(strs...)
	}
}

func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Build(str1, str2, str3)
	}
}

func BenchmarkBuild10values(b *testing.B) {
	strs := []string{"012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678", "012345678"}
	for i := 0; i < b.N; i++ {
		Build(strs...)
	}
}

func BenchmarkRuneToIntByString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RuneToIntByString()
	}
}

func BenchmarkRuneToIntByMinus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RuneToIntByMinus()
	}
}
