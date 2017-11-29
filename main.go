package main

import (
	"fmt"
	"github.com/stormentt/tictactoe/driver"
)

func main() {
	me := uint(257)
	them := uint(72)

	bestMove := uint(0)
	if true {
		bestMove = driver.FindBest(me, them)
	} else {
		bestMove = driver.FindBest(them, me)
	}
	fmt.Println(bestMove)
}
