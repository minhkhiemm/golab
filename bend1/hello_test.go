package bend1

import (
	"fmt"
	"sort"
	"testing"
)

//  go test -bench=. -benchtime 2s -count 2 -benchmem -cpu 4
/*
	While you are running the benchmarks in this book, be sure to remember that benchmarks aren't
	the be-all and end-all for performance result. Benchmarking has both positives and drawbacks:

	The positives of benchmarking are as follows:
		- Surfaces potential problems before they become unwieldy
		- Helps developers have a deeper understanding of their code
		- Can identify potential bottlenecks in the design and data structures and algorithms stages

	The drawbacks of benchmarking are as follows:
		- Needs to be completed on a given cadence fore meaning ful results
		- Data collation can be difficult
		- Does not always yield a meaningful result for the problem at hand
*/
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello Go High Performance")
	}
}

/*
	O(1)

	get element from array
*/
func BenchmarkThreeWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeWords()
	}
}

func BenchmarkTenWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenWords()
	}
}

func BenchmarkSearchTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Search(10, func(i int) bool {
			if i%27 == 0 {
				return true
			}
			return false
		})
	}
}

func BenchmarkSearchHundred(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Search(100, func(i int) bool {
			if i%27 == 0 {
				return true
			}
			return false
		})
	}
}

func BenchmarkSearchThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Search(1000, func(i int) bool {
			if i%27 == 0 {
				return true
			}
			return false
		})
	}
}
