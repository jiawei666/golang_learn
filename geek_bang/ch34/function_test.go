package ch34

import (
	"bytes"
	"testing"
)

func BenchmarkConcatByBuf(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

func BenchmarkConcat(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}
