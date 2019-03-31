package main

import (
	"bytes"
	"runtime"
	"runtime/pprof"
	"testing"
)

func BenchmarkMemstats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ms := &runtime.MemStats{}
		runtime.ReadMemStats(ms)
	}
}

func BenchmarkPProfHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		pprof.Lookup("heap").WriteTo(buf, 1)
	}
}
