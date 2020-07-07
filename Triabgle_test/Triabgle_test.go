package main

import "testing"
// 10 0000 0000
func BenchmarkTriangle(b *testing.B) {
	m, n := 30000, 40000
	ans := 50000
	for i := 0; i < b.N; i++ {
		actual := Triangle(m, n)
		if actual != ans {
			b.Errorf("Triangle(%d, %d);"+
				"got %d; expected %d", m, n, actual, ans)
		}
	}
}
