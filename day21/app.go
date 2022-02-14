package main

import (
	"fmt"
	"time"
)

type Player struct {
	Position int
	Score    int
}

// 3 x 1
// 4 x 3
// 5 x 6
// 6 x 7
// 7 x 6
// 8 x 3
// 9 x 1
var Results = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

var winScore = 21

// Player 1 starting position: 8
// Player 2 starting position: 6
func main() {
	start := time.Now()

	p1 := Player{
		Position: 8,
		Score:    0,
	}

	p2 := Player{
		Position: 6,
		Score:    0,
	}

	p1wins := int64(0)
	p2wins := int64(0)
	deepGame(p1, p2, &p1wins, &p2wins, 1)

	fmt.Println(p1wins, p1wins > p2wins)
	fmt.Println(p2wins, p1wins < p2wins)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func deepGame(p1 Player, p2 Player, p1wins *int64, p2wins *int64, universes int64) {

	for dice1, times1 := range Results {

		nextPosition1 := mod10(p1.Position + dice1)
		nextScore1 := p1.Score + nextPosition1

		if nextScore1 >= winScore {
			*p1wins = *p1wins + universes*int64(times1)
		} else {

			for dice2, times2 := range Results {
				nextPosition2 := mod10(p2.Position + dice2)
				nextScore2 := p2.Score + nextPosition2

				if nextScore2 >= winScore {
					*p2wins = *p2wins + universes*int64(times1)*int64(times2)
				} else {
					nextP1 := Player{
						Position: nextPosition1,
						Score:    nextScore1,
					}
					nextP2 := Player{
						Position: nextPosition2,
						Score:    nextScore2,
					}
					deepGame(nextP1, nextP2, p1wins, p2wins, universes*int64(times1)*int64(times2))
				}
			}

		}
	}

}

// func Roll(p Player, d Dice) (Player, Dice) {
// 	score := mod100(d.Current) + mod100(d.Current+1) + mod100(d.Current+2)

// 	p.Position = mod10(p.Position + score)
// 	p.Score = p.Score + p.Position

// 	d.Rolls = d.Rolls + 3
// 	d.Current = mod100(d.Current + 3)

// 	return p, d
// }

func mod100(n int) int {
	return (n-1)%100 + 1
}

func mod10(n int) int {
	return (n-1)%10 + 1
}

func combinations() {
	scores := map[int]int{}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				s := i + j + k
				scores[s]++
			}
		}
	}

	for k, v := range scores {
		fmt.Println(k, "x", v)
	}

}
