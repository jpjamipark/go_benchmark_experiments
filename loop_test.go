package main

import (
	"strings"
	"testing"
)

func countFor1(s string) int {
	result := 0
	for i := 0; i < len(s); {
		if s[i]&1 == 1 {
			result += 1
		}
		i++
	}

	return result
}

func countFor2(s string) int {
	result := 0
	for i := 0; i < len(s); {
		if s[i]&1 == 1 {
			result += 1
			i++
		} else {
			i++
		}
	}

	return result
}

func genInput(N int) string {
	var sb strings.Builder
	for i := 0; i < N; i++ {
		sb.WriteByte(byte(i % 128))
	}
	return sb.String()
}

const N = 1000000

func BenchmarkFor1(b *testing.B) {
	inp := genInput(N)
	b.SetBytes(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countFor1(inp)
	}
}

func BenchmarkFor2(b *testing.B) {
	inp := genInput(N)
	b.SetBytes(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countFor2(inp)
	}
}



