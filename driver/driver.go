package driver

import "fmt"

var (
	//these are binary representations like
	//111000000
	wins = []uint16{448, 56, 7, 292, 146, 73, 273, 84}
)

func ScoreBoard(me, them uint16) int {
	for i := 0; i < 8; i++ {
		if me&wins[i] == wins[i] {
			return 100
		}
		if them&wins[i] == wins[i] {
			return -100
		}
	}
	return 0
}

func ScoreMove(me, them uint16, depth int, myTurn bool) int {

	//calc the score, if its 100 or -100 the game has been won
	score := ScoreBoard(me, them)
	if score == 100 {
		return score - depth
	} else if score == -100 {
		return score + depth
	}

	// board is completely full
	if me&them == 511 {
		return score
	}

	taken := me | them
	if myTurn {
		best := -1000
		for i := uint16(0); i < 9; i++ {
			if taken&(1<<i) == 0 {
				moveScore := ScoreMove(me|(1<<i), them, depth+1, !myTurn)
				if moveScore > best {
					best = moveScore
				}
			}
		}
		return best
	} else {
		worst := 1000
		for i := uint16(0); i < 9; i++ {
			if taken&(1<<i) == 0 {
				moveScore := ScoreMove(me, them|(1<<i), depth+1, !myTurn)
				if moveScore < worst {
					worst = moveScore
				}
			}
		}
		return worst
	}
}

func FindBest(me, them uint16) uint16 {
	bestSpot := uint16(0)
	bestScore := -1000
	taken := me | them
	//printBoard(me, them)
	//fmt.Println("Starting Board")
	for i := uint16(0); i < 9; i++ {
		if taken&(1<<i) == 0 {
			score := ScoreMove(me|(1<<i), them, 0, false)
			if score > bestScore {
				bestScore = score
				bestSpot = i
			}
		}
	}
	//printBoard(me|(1<<bestSpot), them)
	//fmt.Println("Optimal Move, move weight", bestScore, "new value", me|1<<bestSpot)
	//fmt.Println()
	return bestSpot
}

func printBoard(xs, os uint16) {
	for i := uint16(0); i < 9; i++ {
		if i%3 == 0 {
			fmt.Printf("\n")
		}
		if xs&(1<<i) == (1 << i) {
			fmt.Printf("x")
		} else if os&(1<<i) == 1<<i {
			fmt.Printf("o")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
