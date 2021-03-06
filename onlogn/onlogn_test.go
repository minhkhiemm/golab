package onlogn

import "testing"

/*
O(nlog n)

10: 10s
100: 200s
1000: 3000s


- The average case time complexity for Quicksort
- The average case time complexity for mergesort
- The average case time complexity for Heapsort
- The average case time complexity for Timsort
*/
func benchmarkOnlognLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		OnlognLoop(i)
	}
}

func BenchmarkOnlognLoop10(b *testing.B)   { benchmarkOnlognLoop(10, b) }
func BenchmarkOnlognLoop100(b *testing.B)  { benchmarkOnlognLoop(100, b) }
func BenchmarkOnlognLoop1000(b *testing.B) { benchmarkOnlognLoop(1000, b) }
