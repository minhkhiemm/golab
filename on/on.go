package on

func generateIntSlice(n int) []int {
	x := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x = append(x, i)
	}
	return x
}

func simpleLoopSum(count int) int {
	slice := generateIntSlice(count)
	sum := 0
	for i := 0; i < count; i++ {
		sum += slice[i]
	}
	return sum
}
