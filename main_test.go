package main

import "testing"
import "tictactoe/driver"

func BenchmarkBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		driver.FindBest(8, 1)
	}
}
