package main

import "testing"

// BenchmarkEcho1 test time
// run with `go test -bench=`
func BenchmarkEcho1(b *testing.B) {
	args := []string{"test", "hello", "whatever", "whenever", "gold"}
	for i := 0; i < b.N; i++ {
		Echo1(args)
	}
}

// BenchmarkEcho2 test time
// run with `go test -bench= -benchmem`
func BenchmarkEcho2(b *testing.B) {
	args := []string{"test", "hello", "whatever", "whenever", "gold"}
	for i := 0; i < b.N; i++ {
		Echo2(args)
	}
}

// BenchmarkEcho5 test time
// run with `go test -bench=`
func BenchmarkEcho5(b *testing.B) {
	args := []string{"test", "hello", "whatever", "whenever", "gold"}
	for i := 0; i < b.N; i++ {
		Echo5(args)
	}
}

// BenchmarkEcho6 test time
// run with `go test -bench=`
func BenchmarkEcho6(b *testing.B) {
	args := []string{"test", "hello", "whatever", "whenever", "gold"}
	for i := 0; i < b.N; i++ {
		Echo6(args)
	}
}
