package consensus

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func nBytes(n int) []byte {
	buf := make([]byte, n)
	n, _ = rand.Read(buf)
	return buf[:n]
}

func benchmarkWalDecode(b *testing.B, n int) {
	buf := new(bytes.Buffer)
	enc := NewWALEncoder(buf)

	data := nBytes(n)
	enc.Encode(data)

	encoded := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(encoded)
		dec := NewWALDecoder(buf)
		if _, err := dec.Decode(); err == nil {
		}
	}
	b.ReportAllocs()
}

func BenchmarkWalDecode512B(b *testing.B) {
	benchmarkWalDecode(b, 512)
}

func BenchmarkWalDecode10KB(b *testing.B) {
	benchmarkWalDecode(b, 10*1024)
}
func BenchmarkWalDecode100KB(b *testing.B) {
	benchmarkWalDecode(b, 100*1024)
}
func BenchmarkWalDecode1MB(b *testing.B) {
	benchmarkWalDecode(b, 1024*1024)
}
func BenchmarkWalDecode10MB(b *testing.B) {
	benchmarkWalDecode(b, 10*1024*1024)
}
func BenchmarkWalDecode100MB(b *testing.B) {
	benchmarkWalDecode(b, 100*1024*1024)
}
func BenchmarkWalDecode1GB(b *testing.B) {
	benchmarkWalDecode(b, 1024*1024*1024)
}
