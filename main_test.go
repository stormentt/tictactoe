package main

import (
	"testing"
	"github.com/stormentt/tictactoe/driver"
)

func BenchmarkBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		driver.FindBest(8, 1)
	}
}
